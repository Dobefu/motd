package utils

import (
	"github.com/go-ini/ini"
	"log"
	"runtime"
)

func GetOS() string {
	os := runtime.GOOS

	if os != "linux" {
		return os
	}

	cfg, err := ini.Load("/etc/os-release")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	ConfigParams := make(map[string]string)
	ConfigParams["ID"] = cfg.Section("").Key("ID").String()

	return ConfigParams["ID"]
}
