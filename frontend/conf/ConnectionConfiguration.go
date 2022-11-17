package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func GetApiGateway() string {
	cfg, err := ini.Load("conf/host.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	host := cfg.Section(cfg.Section("").Key("app_mode").String()).Key("host_api_gateway").String()
	port := cfg.Section(cfg.Section("").Key("app_mode").String()).Key("port_api_gateway").String()

	return host + ":" + port
}
