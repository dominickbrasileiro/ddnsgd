package ipv4

import (
	"errors"
	"log"
)

type IPv4Datasource interface {
	GetName() string
	FetchIPv4() (string, bool)
}

type IPv4Repository struct {
	Datasources []IPv4Datasource
	logger      *log.Logger
}

func NewIPv4Repository(datasources []IPv4Datasource, logger *log.Logger) *IPv4Repository {
	return &IPv4Repository{
		Datasources: datasources,
		logger:      logger,
	}
}

func (r *IPv4Repository) FetchIPv4() (string, error) {
	for _, ds := range r.Datasources {
		result, ok := ds.FetchIPv4()

		if ok && ValidateIPv4(result) {
			r.logger.Printf("[IPv4Repository] Fetched IPv4 from %s: %s\n", ds.GetName(), result)
			return result, nil
		}

		r.logger.Println("[IPv4Repository] Failed to fetch IPv4 address from", ds.GetName())
	}

	return "", errors.New("could not fetch ipv4 address")
}
