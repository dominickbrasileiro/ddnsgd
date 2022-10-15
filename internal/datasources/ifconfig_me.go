package datasources

import (
	"io"
	"net/http"

	"github.com/dominickbrasileiro/ddns-google-domains/internal"
)

type IfConfigMeDatasource struct{}

var _ internal.IPv4Datasource = (*IfConfigMeDatasource)(nil)

func NewIfConfigMeDatasource() *IfConfigMeDatasource {
	return &IfConfigMeDatasource{}
}

func (d *IfConfigMeDatasource) FetchIPv4() (string, bool) {
	res, err := http.Get("https://ifconfig.me")

	if err != nil {
		return "", false
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return "", false
	}

	return string(bodyBytes), true
}
