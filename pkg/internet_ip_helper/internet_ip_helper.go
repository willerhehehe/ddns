package internet_ip_helper

import (
	"ddns/pkg/ipv4_http"
	"encoding/json"
	"fmt"
	"net"
)

type BaseInternetIPHelper interface {
	GetInternetIP() (net.IP, error)
}

type InternetIPHelper struct {
}

func (h InternetIPHelper) GetInternetIP() (net.IP, error) {
	urls := []string{"https://api.ipify.org?format=json", "https://jsonip.com", "https://ipinfo.io"}
	type IPResp struct {
		IP net.IP `json:"ip"`
	}
	var ipResp IPResp
	for _, url := range urls {
		resp, err := ipv4_http.Get(url)
		if err != nil {
			continue
		}
		err = json.Unmarshal(resp, &ipResp)
		if err != nil {
			continue
		}
		return ipResp.IP, nil
	}
	return nil, fmt.Errorf("GetInternetIP Error")
}
