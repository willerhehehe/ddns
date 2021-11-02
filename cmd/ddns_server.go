package main

import (
	"ddns/internal/config"
	ddns "ddns/pkg/ddns_server"
	"ddns/pkg/dns_helper/alidns"
	"ddns/pkg/email"
	i "ddns/pkg/internet_ip_helper"
	"fmt"
	"log"
	"net"
)

type Notifier struct {
	email.SimpleEmailClient
}

func (n Notifier) Notify(msg interface{}) {
	t, ok := msg.(string)
	if !ok {
		n.Notify(fmt.Sprintf("Notify Error: msg.(string) error, msg: %v\n", msg))
	}
	err := n.SendMail(n.Username, []string{n.Username}, "DDNS服务通知", t, "text/html")
	if err != nil {
		log.Println(fmt.Sprintf("Notify Error: n.SendMail error, msg: %v\n", msg))
	}
}

func main() {
	conf := config.Conf{}
	conf.ReadConfig()
	server := ddns.DDnsServer{
		IpHelper: i.InternetIPHelper{},
		DnsHelper: alidns.Client{
			AccessKeyId:     conf.AccessKeyId,
			AccessKeySecret: conf.AccessKeySecret,
			RecordId:        conf.RecordId,
		},
		Notifier: Notifier{
			email.SimpleEmailClient{
				Username: conf.Username,
				Password: conf.Password,
				SmtpHost: conf.SmtpHost,
				SmtpPort: conf.SmtpPort,
				Identity: conf.Identity,
			},
		},
		IpChan:  make(chan net.IP),
		MsgChan: make(chan string),
	}
	server.Run()
}
