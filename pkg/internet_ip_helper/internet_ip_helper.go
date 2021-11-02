package internet_ip_helper

import (
	"ddns/pkg/ipv4_http"
	"encoding/json"
	"net"
)

type BaseInternetIPHelper interface {
	GetInternetIP() (net.IP, error)
}

type InternetIPHelper struct {
}

func (h InternetIPHelper) GetInternetIP() (net.IP, error) {
	url := "https://jsonip.com"
	type IPResp struct {
		IP net.IP `json:"ip"`
	}
	var ipResp IPResp
	resp, err := ipv4_http.Get(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &ipResp)
	if err != nil {
		return nil, err
	}
	return ipResp.IP, nil
}
