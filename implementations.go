package main

import (
	"io"
	"net/http"
)

type IpifyDatasource struct{}

var _ IPv4Datasource = (*IpifyDatasource)(nil)

func NewIpifyDatasource() *IpifyDatasource {
	return &IpifyDatasource{}
}

func (d *IpifyDatasource) FetchIPv4() (string, bool) {
	res, err := http.Get("https://api.ipify.org")

	if err != nil {
		return "", false
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return "", false
	}

	return string(bodyBytes), true
}
