package internal

import (
	"log"

	"github.com/dominickbrasileiro/ddnsgd/pkg/googledomains"
	"github.com/dominickbrasileiro/ddnsgd/pkg/ipv4"
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

	ipv4Ch := make(chan string)

	go ipv4.IPv4Polling(config.Interval, logger, ipv4Ch)

	for ip := range ipv4Ch {
		if ip == lastUpdatedIP {
			continue
		}

		err := googledomains.UpdateDDNS(
			ip,
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

		lastUpdatedIP = ip

		logger.Println("[App] Google Domains DDNS updated successfully")
	}
}
