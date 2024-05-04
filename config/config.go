package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		fmt.Println("loade config err")
	}
}

type MailConfig struct {
	HostName string
	Port     string
	Username string
	Password string
}

func Mail() *MailConfig {
	conf := Cfg.Section("mail")
	mc := &MailConfig{
		HostName: func(conf *ini.Section) string { key, _ := conf.GetKey("MAIL_HOST_NAME"); return key.String() }(conf),
		Port:     func(conf *ini.Section) string { key, _ := conf.GetKey("MAIL_PORT"); return key.String() }(conf),
		Username: func(conf *ini.Section) string { key, _ := conf.GetKey("MAIL_USER"); return key.String() }(conf),
		Password: func(conf *ini.Section) string { key, _ := conf.GetKey("MAIL_PASSWORD"); return key.String() }(conf),
	}

	if recover() != nil {
		return &MailConfig{}
	}

	return mc
}
