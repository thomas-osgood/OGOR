package dnsenum

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	errorhandling "github.com/thomas-osgood/OGOR/misc/error-handling"
	"github.com/thomas-osgood/OGOR/misc/generators"
)

// function designed to enumerate a top-level domain for
// subdomains using VHost enumeration. this will use the
// wordlist provided to the enumerator.
func (e *Enumerator) EnumSubdomainsVHOST() (err error) {
	var comms chan string = make(chan string)
	var listgen *generators.WordlistGenerator
	var targetip string
	var tldlen int
	var wgrp *sync.WaitGroup = new(sync.WaitGroup)

	listgen, err = generators.NewWordlistGenerator(e.Wordlist)
	if err != nil {
		return err
	}

	targetip, err = e.getIP()
	if err != nil {
		return err
	}

	tldlen, err = e.getTLDLen()
	if err != nil {
		return err
	}

	listgen.CommsChan = comms

	go listgen.ReadWordlist()

	for i := 0; i < e.threads; i++ {
		wgrp.Add(1)
		go e.vhostWorker(targetip, tldlen, &comms, wgrp)
	}
	wgrp.Wait()

	if len(e.Discovered) < 1 {
		return errors.New("no subdomains found")
	}

	return nil
}

// function designed to enumerate a top-level domain for
// subdomains using GET request enumeration. this will
// use the wordlist provided to the enumerator.
func (e *Enumerator) EnumSubdomainsGET() (err error) {
	var comms chan string = make(chan string)
	var exit bool = false
	var listgen *generators.WordlistGenerator
	var wgrp *sync.WaitGroup = new(sync.WaitGroup)

	listgen, err = generators.NewWordlistGenerator(e.Wordlist)
	if err != nil {
		return err
	}

	if e.proxyscraper != nil {
		e.proxychan = make(chan string)
		go generators.ProxyGenerator(e.proxyscraper, e.proxychan, nil, &exit)
	}

	listgen.CommsChan = comms

	go listgen.ReadWordlist()

	for i := 0; i < e.threads; i++ {
		wgrp.Add(1)
		go e.getWorker(&comms, wgrp)
	}
	wgrp.Wait()

	exit = true

	if len(e.Discovered) < 1 {
		return errors.New("no subdomains found")
	}

	return nil
}

// function designed to grab the return length of a
// GET request to the top-level domain.
func (e *Enumerator) GetTLDLength() (err error) {
	var bodycontent []byte
	var resp *http.Response
	var targeturl string

	targeturl = fmt.Sprintf("http://%s", e.TLD)

	resp, err = e.Client.Get(targeturl)

	// HTTP HEAD request returned error, attempt with HTTPS.
	if (err != nil) || (resp.StatusCode >= http.StatusBadRequest) {
		targeturl = fmt.Sprintf("https://%s", e.TLD)
		resp, err = e.Client.Get(targeturl)
		if err != nil {
			e.TLDLength = -1
			return err
		}

		// HTTPS HEAD returned error, exit with failure.
		if resp.StatusCode >= http.StatusBadRequest {
			return errors.New(fmt.Sprintf("\"%s\" unreachable by HEAD", e.TLD))
		}

		if (resp.StatusCode == http.StatusFound) || (resp.StatusCode == http.StatusMovedPermanently) {
			targeturl = resp.Header.Get("location")

			resp, err = e.Client.Get(targeturl)
			if err != nil {
				return err
			}
		}
	}

	// one of the above HEAD requests succeeded. get length.
	defer resp.Body.Close()
	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	e.TLDLength = len(bodycontent)

	return nil
}

// function designed to use a google dork to attempt to locate
// subdomains of the given domain.
func (e *Enumerator) GoogleDork() (err error) {
	var bodydata []byte
	var client http.Client = http.Client{}
	var dork string = fmt.Sprintf(`site:*.%s`, e.TLD)
	var match []byte
	var matches [][]byte
	var matchstr string
	var pattern string = fmt.Sprintf(`a href="(\/url\?q=)?http(s)?:\/\/[a-zA-Z0-9]+\.%s`, strings.ReplaceAll(e.TLD, ".", `\.`))
	var re *regexp.Regexp
	var req *http.Request
	var resp *http.Response
	var targeturl string

	targeturl = fmt.Sprintf("https://www.google.com/search?q=%s", url.PathEscape(dork))

	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return err
	}

	if e.display {
		e.printer.SysMsgNB("dorking google ...")
	}

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New(fmt.Sprintf("dork returned bad status code (%s)", resp.Status))
	}

	bodydata, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	re, err = regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if e.display {
		e.printer.SysMsgNB("searching response for subdomains ...")
	}

	matches = re.FindAll(bodydata, -1)
	if matches == nil {
		return errors.New("no results discovered in dork")
	}

	for _, match = range matches {
		matchstr = strings.Split(strings.Split(strings.Split(string(match), "//")[1], "/")[0], "?")[0]
		err = e.addSubdomain(matchstr)
	}

	if e.display {
		for _, subdomain := range e.Discovered {
			e.printer.SucMsg(subdomain)
		}
	}

	return nil
}

// function designed to see if the TLD provided is reachable
// using HEAD and is, therefore, a valid site or if the domain
// is comprised enitrely of subdomains (if any).
func (e *Enumerator) TestTLD() (err error) {
	err = e.TestSubdomainHead(e.TLD, false)
	if err != nil {
		err = e.TestSubdomainHead(e.TLD, true)
		if err != nil {
			return err
		}
	}
	return nil
}

// function designed to test a subdomain to see if it returns without
// error when a HEAD request is made against it.
//
// note: the subdomain must be <subdomain>.<tld> (ex: test.example.com)
func (e *Enumerator) TestSubdomainHead(subdomain string, https bool) (err error) {
	var ok bool = true
	var proxystr string
	var proxyurl *url.URL
	var resp *http.Response
	var targeturl string

	if https {
		targeturl = fmt.Sprintf("https://%s", subdomain)
	} else {
		targeturl = fmt.Sprintf("http://%s", subdomain)
	}

	// if a proxyscraper object is present, use proxies to
	// anonymize the requests.
	if e.proxyscraper != nil {
		proxystr, ok = <-e.proxychan
		if !ok {
			return errors.New("error communicating with proxy channel")
		}

		proxyurl, err = url.Parse(proxystr)
		if err != nil {
			return err
		}

		e.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyurl)}
	}

	resp, err = e.Client.Head(targeturl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// function designed to test a subdomain to see if it returns without
// error when the HOST header is manipulated.
//
// note: the subdomain must be <subdomain>.<tld> (ex: test.example.com)
func (e *Enumerator) TestSubdomainHost(subdomain string, targetip string, tldlen int, https bool) (err error) {
	var bodydata []byte
	var req *http.Request
	var resp *http.Response
	var scheme string
	var targeturl string

	if https {
		scheme = "https"
	} else {
		scheme = "http"
	}

	targeturl = fmt.Sprintf("%s://%s", scheme, targetip)

	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return err
	}

	req.Host = subdomain

	resp, err = e.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if (resp.StatusCode >= 400) && !((resp.StatusCode == http.StatusUnauthorized) || (resp.StatusCode == http.StatusForbidden)) {
		return errors.New("invalid subdomain")
	}

	bodydata, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(bodydata) == tldlen {
		return errors.New(fmt.Sprintf("length matches tld (%d)", len(bodydata)))
	}

	return nil
}

// function designed to add a subdomain to the Discovered slice. this
// will only add the subdomain if it does not already exist within
// the slice.
func (e *Enumerator) addSubdomain(subdomain string) (err error) {

	err = e.subdomainExists(subdomain)
	if (err != nil) && !(errors.Is(err, &errorhandling.NotFoundError{})) {
		return err
	} else if err == nil {
		return nil
	}

	e.Discovered = append(e.Discovered, subdomain)
	return nil
}

// function designed to calculate a delay based on the max allowed
// delay of the enumerator.
func (e *Enumerator) calculateDelay() (delay time.Duration, err error) {
	if e.delay > 0 {
		rand.Seed(time.Now().UnixMicro())
		delay = time.Duration(rand.Intn(e.delay)) * time.Millisecond
	} else {
		delay = 0
	}
	return delay, nil
}

// function to get the IP address of the target domain. this will
// be used in VHOST enumeration.
func (e *Enumerator) getIP() (ip string, err error) {
	var curip net.IP
	var curip4 net.IP
	var ips []net.IP

	ips, err = net.LookupIP(e.TLD)
	if err != nil {
		return "", err
	}

	for _, curip = range ips {
		if curip4 = curip.To4(); curip4 != nil {
			ip = curip4.String()
			break
		}
	}

	return ip, nil
}

// function designed to get the page length of the TLD that is being
// enumerated. this will be used to determine if a VHOST is valid or
// a false positive.
func (e *Enumerator) getTLDLen() (pagelen int, err error) {
	var bodydata []byte
	var req *http.Request
	var resp *http.Response
	var scheme string
	var targeturl string

	if e.https {
		scheme = "https"
	} else {
		scheme = "http"
	}

	targeturl = fmt.Sprintf("%s://%s", scheme, e.TLD)

	req, err = http.NewRequest(http.MethodGet, targeturl, nil)
	if err != nil {
		return -1, err
	}

	resp, err = e.Client.Do(req)
	if err != nil {
		return -1, err
	}

	bodydata, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	pagelen = len(bodydata)

	return pagelen, nil
}

// function designed to be a worker thread for the GET enumeration.
// this will return nil if no errors occur during testing, otherwise
// and error is returned.
func (e *Enumerator) getWorker(comms *chan string, wgrp *sync.WaitGroup) (err error) {
	var subdomain string
	var target string
	var delay time.Duration

	delay, err = e.calculateDelay()
	if err != nil {
		wgrp.Done()
		return err
	}

	for subdomain = range *comms {
		target = fmt.Sprintf("%s.%s", subdomain, e.TLD)

		if e.display {
			e.printer.SysMsgNB(fmt.Sprintf("testing %s", target))
		}

		err = e.TestSubdomainHead(target, e.https)
		if err != nil {
			time.Sleep(delay)
			continue
		}
		err = e.addSubdomain(target)

		if e.display {
			if err != nil {
				e.printer.ErrMsg(err.Error())
			}
			e.printer.SucMsg(target)
		}

		time.Sleep(delay)
	}

	wgrp.Done()
	return nil
}

// function designed to check if a given subdomain exists within
// the Discovered slice of the Enumerator. if the subdomain does
// not exist, a NotFoundError is returned.
func (e *Enumerator) subdomainExists(subdomain string) (err error) {
	var current string
	var found bool = false

	for _, current = range e.Discovered {
		if current == subdomain {
			found = true
			break
		}
	}

	if !found {
		return &errorhandling.NotFoundError{}
	}

	return nil
}

// function designed to be a worker thread for the VHost enumeration.
// this will return nil if no errors occur during testing, otherwise
// and error is returned.
func (e *Enumerator) vhostWorker(targetip string, tldlen int, comms *chan string, wgrp *sync.WaitGroup) (err error) {
	var delay time.Duration
	var subdomain string
	var target string

	delay, err = e.calculateDelay()
	if err != nil {
		wgrp.Done()
		return err
	}

	for subdomain = range *comms {
		target = fmt.Sprintf("%s.%s", subdomain, e.TLD)

		if e.display {
			e.printer.SysMsgNB(fmt.Sprintf("testing %s", target))
		}

		err = e.TestSubdomainHost(target, targetip, tldlen, e.https)
		if err != nil {
			time.Sleep(delay)
			continue
		}
		err = e.addSubdomain(target)

		if e.display {
			if err != nil {
				e.printer.ErrMsg(err.Error())
			}
			e.printer.SucMsg(target)
		}

		time.Sleep(delay)
	}

	wgrp.Done()
	return nil
}
