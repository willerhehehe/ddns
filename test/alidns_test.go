package test

import (
	"ddns/pkg/dns_helper/alidns"
	"testing"
)

func TestAliDNSGet(t *testing.T) {
	client := alidns.Client{
		AccessKeyId:     "",
		AccessKeySecret: "",
		RecordId:        "",
	}
	ip, err := client.GetDnsIP()
	if err != nil {
		t.Log(err)
	}
	t.Log(ip)
	t.Log(ip.String())

}
