package ipv4

import (
	"io"
	"net/http"
)

type IpInfoDatasource struct{}

var _ IPv4Datasource = (*IpInfoDatasource)(nil)

func NewIpInfoDatasource() *IpInfoDatasource {
	return &IpInfoDatasource{}
}

func (d *IpInfoDatasource) FetchIPv4() (string, bool) {
	res, err := http.Get("https://ipinfo.io/ip")

	if err != nil {
		return "", false
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return "", false
	}

	return string(bodyBytes), true
}
