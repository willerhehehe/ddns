package config

import (
	"embed"
	"gopkg.in/yaml.v3"
	"log"
)

//go:embed config.yaml
var config embed.FS

type Conf struct {
	AccessKeyId     string `yaml:"AccessKeyId"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	RecordId        string `yaml:"RecordId"`
	Username        string `yaml:"Username"`
	Password        string `yaml:"Password"`
	SmtpHost        string `yaml:"SmtpHost"`
	SmtpPort        string `yaml:"SmtpPort"`
	Identity        string `yaml:"Identity"`
}

func (conf *Conf) ReadConfig() *Conf {
	yamlFile, err := config.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
