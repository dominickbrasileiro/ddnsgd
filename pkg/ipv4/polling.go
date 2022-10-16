package ipv4

import (
	"log"
	"time"
)

func IPv4Polling(interval int, logger *log.Logger, resChannel chan string) {
	datasources := []IPv4Datasource{
		NewIpifyDatasource(),
		NewIpInfoDatasource(),
		NewIfConfigMeDatasource(),
	}

	repository := NewIPv4Repository(datasources, logger)
	ticker := time.NewTicker(time.Second * time.Duration(interval))

	for range ticker.C {
		ipv4, err := repository.FetchIPv4()

		if err != nil {
			logger.Println("[App] Failed to fetch IPv4 address")
			continue
		}

		resChannel <- ipv4
	}
}
