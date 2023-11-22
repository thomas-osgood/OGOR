package forwardproxy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/thomas-osgood/OGOR/networking/apis"
)

// async function designed to transfer connections made
// while handling the CONNECT request.
func (f *Forwarder) makeTunnelConnection(source io.ReadCloser, dest io.WriteCloser) {
	defer source.Close()
	defer dest.Close()
	io.Copy(dest, source)
}

// function designed to display the request information
// being passed through the proxy. this will show the
// method, headers, body, etc.
func (f *Forwarder) displayRequestInfo(r *http.Request) (err error) {
	var headers map[string][]string = r.Header
	var headerKey string
	var headerVals []string
	var headerVal string
	var method string = r.Method
	var path string = r.URL.Path
	var scheme string = r.URL.Scheme
	var url string = r.URL.Host

	fmt.Printf("%s\n", strings.Repeat("-", 60))

	fmt.Printf("URL: %s://%s%s\n", scheme, url, path)
	fmt.Printf("Method: %s\n", method)

	for headerKey, headerVals = range headers {
		fmt.Printf("%s:\n", headerKey)
		for _, headerVal = range headerVals {
			fmt.Printf("\t%s\n", headerVal)
		}
	}

	fmt.Printf("%s\n", strings.Repeat("-", 60))

	return nil
}

// function designed to handle when a client sends a CONNECT
// request to the target. this will handle it gracefully so
// no error occurs. after this is processed, the GET, POST, etc
// requests can be properly handled by the transmitRequest func.
//
// ref: https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c
func (f *Forwarder) openTunnel(w http.ResponseWriter, r *http.Request) (err error) {
	var client_conn net.Conn
	var dest_conn net.Conn

	dest_conn, err = net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	client_conn, _, err = hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	go f.makeTunnelConnection(client_conn, dest_conn)
	go f.makeTunnelConnection(dest_conn, client_conn)

	log.Printf("\"%s\" <----> \"%s%s\"", r.RemoteAddr, r.URL.Host, r.URL.Path)

	return nil
}

// function designed to transmit the request to the
// target destination.
//
// ref: https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-1-http-proxies/
// ref: https://reintech.io/blog/creating-simple-proxy-server-with-go
// ref: https://gist.github.com/afdalwahyu/d2d8374879ecdeea2bca92c97056efee
func (f *Forwarder) transmitRequest(w http.ResponseWriter, r *http.Request) (err error) {
	var bodyData []byte
	var forwardRequest *http.Request
	var headerName string
	var headerValue string
	var headerValues []string
	var proxyResp *http.Response
	var targetUrl *url.URL

	if r.Method == http.MethodConnect {
		return f.openTunnel(w, r)
	}

	targetUrl = r.URL

	// read original request body. this will be used
	// when creating a new request object to transmit
	// to the target.
	bodyData, err = io.ReadAll(r.Body)
	if err != nil {
		log.Printf(err.Error())
	}

	err = f.displayRequestInfo(r)
	if err != nil {
		return apis.ReturnErrorJSON(&w, 500, err.Error())
	}

	// make a new request object to send to the target.
	forwardRequest, err = http.NewRequest(r.Method, targetUrl.String(), bytes.NewReader(bodyData))
	if err != nil {
		log.Printf("new request: %s\n", err.Error())
		return apis.ReturnErrorJSON(&w, 500, err.Error())
	}

	// copy headers to the new request object.
	for headerName, headerValues = range r.Header {
		for _, headerValue = range headerValues {
			forwardRequest.Header.Add(headerName, headerValue)
		}
	}

	proxyResp, err = f.forwardTransport.RoundTrip(forwardRequest)
	if err != nil {
		log.Printf("roundtrip: %s\n", err.Error())
		return apis.ReturnErrorJSON(&w, 500, err.Error())
	}
	defer proxyResp.Body.Close()

	for headerName, headerValues = range proxyResp.Header {
		for _, headerValue = range headerValues {
			w.Header().Add(headerName, headerValue)
		}
	}
	w.WriteHeader(proxyResp.StatusCode)

	_, err = io.Copy(w, proxyResp.Body)
	if err != nil {
		log.Printf("error copying body to writer: %s\n", err.Error())
	}

	return nil
}

// function designed to start the forwarder server
// and begin forwarding traffic.
func (f *Forwarder) Serve() (err error) {

	// set the server address using the portno set
	// during forwarder initialization.
	f.server.Handler = apis.MakeHTTPHandleFunc(f.transmitRequest)

	log.Printf("serving on \"%s\"\n", f.server.Addr)

	return f.server.ListenAndServe()
}
