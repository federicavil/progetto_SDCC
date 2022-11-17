package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func GetConnectionConf(service string) string {
	cfg, err := ini.Load("conf/host.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	host := cfg.Section(cfg.Section("").Key("app_mode").String()).Key("host_" + service).String()
	port := cfg.Section(cfg.Section("").Key("app_mode").String()).Key("port_" + service).String()
	fmt.Println("Connection conf: " + host + ":" + port)
	return host + ":" + port

}
