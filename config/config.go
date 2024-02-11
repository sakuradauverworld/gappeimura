package config

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	Dbname    string
	LogFile   string
	Static    string
}

var Config ConfigList

func init() {
	fmt.Println("aaa")
	LoadConfig()

}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Dbname:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
