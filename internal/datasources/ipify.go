package datasources

import (
	"io"
	"net/http"

	"github.com/dominickbrasileiro/ddns-google-domains/internal"
)

type IpifyDatasource struct{}

var _ internal.IPv4Datasource = (*IpifyDatasource)(nil)

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
