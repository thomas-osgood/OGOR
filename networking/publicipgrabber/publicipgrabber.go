// package designed to query api.whatismyip.com for public
// IP address information. this package contains the structs
// defining the JSON responses expected and the objects that
// will utilize them to pull down information.
package publicipgrabber

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// function designed to query api.whatismyip.com and
// pull down the public IP address information for the
// machine executing the program.
func (ipg *PublicIPGrabber) GetMyIPInformation() (err error) {
	const acceptHeader string = "*/*"
	const acceptEncoding string = "json"
	const origin string = "https://www.whatismyip.com"
	const referer string = "https://www.whatismyip.com"
	const secgpc string = "1"
	const useragent string = "grabber"

	var bodycontent []byte
	var client *http.Client = http.DefaultClient
	var jsonResp PublicIPInfo = PublicIPInfo{}
	var req *http.Request
	var resp *http.Response
	var targeturl string = fmt.Sprintf("%s/wimi.php", BASE_URL)

	req, err = http.NewRequest(http.MethodPost, targeturl, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept-Encoding", acceptEncoding)
	req.Header.Set("Accept", acceptHeader)
	req.Header.Set("Origin", origin)
	req.Header.Set("Referer", referer)
	req.Header.Set("Sec-GPC", secgpc)
	req.Header.Set("User-Agent", useragent)

	resp, err = client.Do(req)
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
