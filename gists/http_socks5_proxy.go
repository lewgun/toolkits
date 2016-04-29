package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
)

const (
	PROXY_ADDR = "127.0.0.1:9050"
	URL        = "http://skunksworkedp2cg.onion/sites.html"
)

func main() {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	// create a request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		os.Exit(2)
	}
	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)
		os.Exit(3)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)
		os.Exit(4)
	}
	fmt.Println(string(b))
}
