package internal

import (
	"log"
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

var lastUpdatedIP string

func Run(config *AppConfig, logger *log.Logger) {
	logger.Printf("[App] Starting DDNS %s | Interval: %d seconds\n", config.Hostname, config.Interval)

	datasources := []ipv4.IPv4Datasource{
		ipv4.NewIpifyDatasource(),
		ipv4.NewIpInfoDatasource(),
		ipv4.NewIfConfigMeDatasource(),
	}

	repository := ipv4.NewIPv4Repository(datasources, logger)

	ticker := time.NewTicker(time.Second * time.Duration(config.Interval))

	for range ticker.C {
		ipv4, err := repository.FetchIPv4()

		if err != nil {
			logger.Println("[App] Failed to fetch IPv4 address")
			continue
		}

		if ipv4 == lastUpdatedIP {
			continue
		}

		err = googledomains.UpdateDDNS(
			ipv4,
			config.Username,
			config.Password,
			config.Hostname,
		)

		if err != nil {
			re, ok := err.(googledomains.ApiError)

			if ok {
				logger.Printf("[App] Failed to update Google Domains DDNS (%s)\n", re.Code)
			} else {
				logger.Printf("[App] Failed to update Google Domains DDNS\n")
			}

			continue
		}

		lastUpdatedIP = ipv4
		logger.Println("[App] Google Domains DDNS updated successfully")
	}
}
