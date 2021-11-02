package test

import (
	ddns "ddns/pkg/ddns_server"
	i "ddns/pkg/internet_ip_helper"
	"fmt"
	"net"
	"testing"
)

type testDnsHelper struct {
	t *testing.T
}

func (d testDnsHelper) GetDnsIP() (net.IP, error) {
	ip := net.ParseIP("27.115.15.110")
	d.t.Log(fmt.Sprintf("获取到IP %v\n", ip))
	return ip, nil
}

func (d testDnsHelper) UpdateDns(ip net.IP) error {
	d.t.Log(fmt.Sprintf("更新IP成功 %v\n", ip))
	return nil
}

type testNotifier struct {
	t *testing.T
}

func (n testNotifier) Notify(msg interface{}) {
	n.t.Log(fmt.Sprintf("notify: %v\n", msg))
}

func TestDDnsServer(t *testing.T) {
	server := ddns.DDnsServer{
		IpHelper:  i.InternetIPHelper{},
		DnsHelper: testDnsHelper{t},
		Notifier:  testNotifier{t},
		IpChan:    make(chan net.IP),
		MsgChan:   make(chan string),
	}
	server.Run()
}
