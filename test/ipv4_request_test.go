package test

import (
	"ddns/pkg/ipv4_http"
	"log"
	"testing"
)

func TestNormalGet(t *testing.T) {
	url := "https://jsonip.com"
	resp, err := ipv4_http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	t.Log(string(resp))
}
