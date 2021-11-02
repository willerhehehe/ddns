package alidns

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"net"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	RecordId        string
}

func (c Client) GetDnsIP() (net.IP, error) {
	client, _err := CreateClient(tea.String(c.AccessKeyId), tea.String(c.AccessKeySecret))
	if _err != nil {
		return nil, _err
	}
	describeDomainRecordInfoRequest := &alidns20150109.DescribeDomainRecordInfoRequest{RecordId: &c.RecordId}
	// 复制代码运行请自行打印 API 的返回值
	res, _err := client.DescribeDomainRecordInfo(describeDomainRecordInfoRequest)
	if _err != nil {
		return nil, _err
	}
	ip := net.ParseIP(*res.Body.Value)
	return ip, _err
}

func (c Client) UpdateDns(ip net.IP) error {
	value := ip.String()
	client, _err := CreateClient(tea.String(c.AccessKeyId), tea.String(c.AccessKeySecret))
	if _err != nil {
		return _err
	}
	rr := "@"
	domainType := "A"
	var ttl int64
	ttl = 600
	recordId := c.RecordId
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		Lang:     tea.String("golang"),
		RecordId: &recordId,
		RR:       &rr,
		Type:     &domainType,
		Value:    &value,
		TTL:      &ttl,
	}
	// 复制代码运行请自行打印 API 的返回值
	_, _err = client.UpdateDomainRecord(updateDomainRecordRequest)
	if _err != nil {
		return _err
	}
	return _err
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dns.aliyuncs.com")
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}
