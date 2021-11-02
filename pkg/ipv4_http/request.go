package ipv4_http

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var RoundTripper http.RoundTripper = &http.Transport{
	Proxy:                 http.ProxyFromEnvironment,
	DialContext:           dialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

func dialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return (&net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 5 * time.Second,
	}).DialContext(ctx, "tcp4", address)
}

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: RoundTripper,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
