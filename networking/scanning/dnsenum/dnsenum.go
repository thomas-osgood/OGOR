package dnsenum

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"

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

	for i := 0; i < 10; i++ {
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
	var listgen *generators.WordlistGenerator
	var wgrp *sync.WaitGroup = new(sync.WaitGroup)

	listgen, err = generators.NewWordlistGenerator(e.Wordlist)
	if err != nil {
		return err
	}

	listgen.CommsChan = comms

	go listgen.ReadWordlist()

	for i := 0; i < 10; i++ {
		wgrp.Add(1)
		go e.getWorker(&comms, wgrp)
	}
	wgrp.Wait()

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
	var resp *http.Response
	var targeturl string

	if https {
		targeturl = fmt.Sprintf("https://%s", subdomain)
	} else {
		targeturl = fmt.Sprintf("http://%s", subdomain)
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
		e.printer.ErrMsg(err.Error())
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
		e.Discovered = append(e.Discovered, target)

		if e.display {
			e.printer.SucMsg(target)
		}

		time.Sleep(delay)
	}

	wgrp.Done()
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
		e.Discovered = append(e.Discovered, target)

		if e.display {
			e.printer.SucMsg(target)
		}

		time.Sleep(delay)
	}

	wgrp.Done()
	return nil
}
