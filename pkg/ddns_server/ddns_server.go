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
	Notify(interface{}) error
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

func (s *DDnsServer) RetryNotify(msg string, retryTimes int, duration time.Duration) {
	for i := 0; i < retryTimes; i++ {
		err := s.Notifier.Notify(msg)
		if err != nil {
			time.Sleep(duration)
			continue
		} else {
			break
		}
	}
}

func (s *DDnsServer) Run() {
	s.RetryNotify(
		fmt.Sprintf("Start DDNS Server, Time: %v\n", time.Now()),
		10,
		time.Second*10,
	)
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
			_ = s.Notifier.Notify(msg)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-c
		switch sig {
		case syscall.SIGQUIT:
			_ = s.Notifier.Notify(fmt.Sprintf("DDnsServer stop by SIGQUIT"))
			return
		case syscall.SIGTERM:
			_ = s.Notifier.Notify(fmt.Sprintf("DDnsServer stop by SIGTERM"))
			return
		case syscall.SIGINT:
			_ = s.Notifier.Notify(fmt.Sprintf("DDnsServer stop by SIGINT"))
			return
		case syscall.SIGHUP:
		default:
			_ = s.Notifier.Notify(fmt.Sprintf("DDnsServer stop by DEFAULT"))
			return
		}
	}
}
