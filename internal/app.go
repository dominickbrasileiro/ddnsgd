package internal

import (
	"time"

	"github.com/dominickbrasileiro/ddns-google-domains/pkg/googledomains"
	"github.com/dominickbrasileiro/ddns-google-domains/pkg/ipv4"
)

type AppConfig struct {
	Interval int
	Username string
	Password string
	Hostname string
}

func Run(c *AppConfig) {
	datasources := []ipv4.IPv4Datasource{
		ipv4.NewIpifyDatasource(),
		ipv4.NewIpInfoDatasource(),
		ipv4.NewIfConfigMeDatasource(),
	}

	repository := ipv4.NewIPv4Repository(datasources)

	ticker := time.NewTicker(time.Second * time.Duration(c.Interval))

	for range ticker.C {
		ipv4, err := repository.FetchIPv4()

		if err != nil {
			continue
		}

		err = googledomains.UpdateDDNS(ipv4, c.Username, c.Password, c.Hostname)

		if err != nil {
			continue
		}
	}
}
