package ddnsserver

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type IPHelper interface {
	GetInternetIP() (net.IP, error)
}

type DnsHelper interface {
	GetDnsIP() (net.IP, error)
	UpdateDns(net.IP) error
}

type Notifier interface {
	Notify(interface{})
}

type DDnsServer struct {
	IpHelper  IPHelper
	DnsHelper DnsHelper
	Notifier  Notifier
	IpChan    chan net.IP
	MsgChan   chan string
}

func (s *DDnsServer) compareIP() {
	intIP, err := s.IpHelper.GetInternetIP()
	if err != nil {
		s.MsgChan <- fmt.Sprintf("GetInternetIP Error: %v\n", err)
		return
	}
	dnsIP, err := s.DnsHelper.GetDnsIP()
	if err != nil {
		s.MsgChan <- fmt.Sprintf("GetDnsID Error: %v\n", err)
		return
	}
	if !intIP.Equal(dnsIP) {
		s.IpChan <- intIP
		s.MsgChan <- fmt.Sprintf("IP变动提醒 IntIP：%v DnsID: %v\n", intIP, dnsIP)
	}
}

func (s *DDnsServer) Run() {
	s.Notifier.Notify(fmt.Sprintf("Start DDNS Server, Time: %v\n", time.Now()))
	go func() {
		for {
			s.compareIP()
			time.Sleep(time.Second * 5)
		}
	}()
	go func() {
		for {
			ip := <-s.IpChan
			err := s.DnsHelper.UpdateDns(ip)
			if err != nil {
				s.MsgChan <- fmt.Sprintf("UpdateDns Error %v\nDnsIP %v\n", err, ip)
			}
		}
	}()
	go func() {
		for {
			msg := <-s.MsgChan
			s.Notifier.Notify(msg)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-c
		s.Notifier.Notify(fmt.Sprintf("Get Signal: %v\n", sig))
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.Notifier.Notify(fmt.Sprintf("DDnsServer stop"))
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
