package dnsenum

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/thomas-osgood/OGOR/misc/generators"
)

// function designed to enumerate a top-level domain for
// subdomains using Virtual Host enumeration. this will
// use the wordlist provided to the enumerator.
func (e *Enumerator) EnumSubdomainsVHOST() (err error) {
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
		go e.vhostWorker(&comms, wgrp)
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

// function designed to test a subdomain to see if returns without
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

// function designed to be a worker thread for the VHost enumeration.
// this will return nil if no errors occur during testing, otherwise
// and error is returned.
func (e *Enumerator) vhostWorker(comms *chan string, wgrp *sync.WaitGroup) (err error) {
	var subdomain string
	var target string

	for subdomain = range *comms {
		target = fmt.Sprintf("%s.%s", subdomain, e.TLD)
		e.printer.SysMsgNB(fmt.Sprintf("testing %s", target))
		err = e.TestSubdomainHead(target, e.https)
		if err != nil {
			continue
		}
		e.Discovered = append(e.Discovered, target)
		e.printer.SucMsg(target)
	}

	wgrp.Done()
	return nil
}
