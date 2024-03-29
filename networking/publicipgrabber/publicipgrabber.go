// package designed to query api.whatismyip.com for public
// IP address information. this package contains the structs
// defining the JSON responses expected and the objects that
// will utilize them to pull down information.
package publicipgrabber

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

// function designed to query api.whatismyip.com and get information
// related to a given IP address.
func (ipg *PublicIPGrabber) GetIPInformation(targetip string) (ipinformation *AppResponse, err error) {
	var bodycontent []byte
	var databuffer *bytes.Buffer = new(bytes.Buffer)
	var datawriter *multipart.Writer = multipart.NewWriter(databuffer)
	var errorstruct ErrorResponse = ErrorResponse{}
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/app.php", BASE_URL)

	ipinformation = new(AppResponse)

	// build required multipart data. this specifies the action being
	// performed along with the target to perform the action on.
	err = ipg.setLookupMultipartData(datawriter, "ip-lookup", "ip", targetip)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest(http.MethodPost, targeturl, databuffer)
	if err != nil {
		return nil, err
	}

	// set content-type header specifying multipart data type.
	req.Header.Set("Content-Type", datawriter.FormDataContentType())

	err = ipg.setRequestHeaders(req)
	if err != nil {
		return nil, err
	}

	resp, err = ipg.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(fmt.Sprintf("error getting ip information: %s", resp.Status))
	}

	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodycontent, &errorstruct)
	if err != nil {
		return nil, err
	} else if len(errorstruct.Error) > 0 {
		return nil, errors.New(errorstruct.Error)
	}

	err = json.Unmarshal(bodycontent, ipinformation)
	if err != nil {
		return nil, err
	}

	return ipinformation, nil
}

// function designed to query api.whatismyip.com and
// pull down the public IP address information for the
// machine executing the program.
func (ipg *PublicIPGrabber) GetMyIPInformation() (err error) {

	var bodycontent []byte
	var jsonResp PublicIPInfo = PublicIPInfo{}
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/wimi.php", BASE_URL)

	req, err = http.NewRequest(http.MethodPost, targeturl, nil)
	if err != nil {
		return err
	}

	err = ipg.setRequestHeaders(req)
	if err != nil {
		return err
	}

	resp, err = ipg.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return errors.New(fmt.Sprintf("error contacting site: %s", resp.Status))
	}

	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodycontent, &jsonResp)
	if err != nil {
		return err
	}

	if jsonResp.Ip == "" {
		return errors.New("unable to find public IP")
	}

	ipg.PublicIP = jsonResp

	return nil
}

// function designed to contact api.whatismyip.com and pull down
// the IP address attached to a given URL.
func (ipg *PublicIPGrabber) GetUrlIP(target string) (arecords *DnsResponse, err error) {
	var bodycontent []byte
	var databuffer *bytes.Buffer = new(bytes.Buffer)
	var datawriter *multipart.Writer = multipart.NewWriter(databuffer)
	var errorstruct ErrorResponse = ErrorResponse{}
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/app.php", BASE_URL)

	arecords = new(DnsResponse)

	// build required multipart data. this specifies the action being
	// performed along with the target to perform the action on.
	err = ipg.setLookupMultipartData(datawriter, "dns-lookup", "url", target)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest(http.MethodPost, targeturl, databuffer)
	if err != nil {
		return nil, err
	}

	// set content-type header specifying multipart data type.
	req.Header.Set("Content-Type", datawriter.FormDataContentType())

	err = ipg.setRequestHeaders(req)
	if err != nil {
		return nil, err
	}

	resp, err = ipg.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(fmt.Sprintf("error getting DNS info: %s", resp.Status))
	}

	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodycontent, &errorstruct)
	if err != nil {
		return nil, err
	} else if len(errorstruct.Error) > 0 {
		return nil, errors.New(errorstruct.Error)
	}

	err = json.Unmarshal(bodycontent, arecords)
	if err != nil {
		return nil, err
	}

	return arecords, nil
}

// function designed to setup the multipart data used in
// a call to api.whatismyip.com when querying information
// related to a given IP address.
func (ipg *PublicIPGrabber) setLookupMultipartData(datawriter *multipart.Writer, action string, bodyname string, targetip string) (err error) {
	defer datawriter.Close()

	var header textproto.MIMEHeader = make(textproto.MIMEHeader)
	var part io.Writer

	header.Set("Content-Disposition", "form-data; name=\"action\"")

	part, err = datawriter.CreatePart(header)
	if err != nil {
		return err
	}
	part.Write([]byte(action))

	header.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"%s\"", bodyname))
	part, err = datawriter.CreatePart(header)
	if err != nil {
		return err
	}
	part.Write([]byte(targetip))

	return nil
}

// function designed to set the proper header values
// for a request to api.whatismyip.com. these headers
// will be standard for each api.whatismyip.com request.
func (ipg *PublicIPGrabber) setRequestHeaders(req *http.Request) (err error) {
	const acceptHeader string = "*/*"
	const acceptEncoding string = "json"
	const origin string = "https://www.whatismyip.com"
	const referer string = "https://www.whatismyip.com"
	const secgpc string = "1"
	const useragent string = "grabber"

	req.Header.Set("Accept-Encoding", acceptEncoding)
	req.Header.Set("Accept", acceptHeader)
	req.Header.Set("Origin", origin)
	req.Header.Set("Referer", referer)
	req.Header.Set("Sec-GPC", secgpc)
	req.Header.Set("User-Agent", useragent)

	return nil
}
