package main

import "errors"

type IPv4Datasource interface {
	FetchIPv4() (string, bool)
}

type IPv4Repository struct {
	Datasources []IPv4Datasource
}

func NewIPv4Repository(datasources []IPv4Datasource) *IPv4Repository {
	return &IPv4Repository{
		Datasources: datasources,
	}
}

func (r *IPv4Repository) FetchIPv4() (string, error) {
	for _, ds := range r.Datasources {
		result, ok := ds.FetchIPv4()

		if ok {
			return result, nil
		}
	}

	return "", errors.New("could not fetch ipv4 address")
}
