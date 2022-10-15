package googledomains

import (
	"fmt"
	"net/http"
)

func UpdateDDNS(ip string, username string, password string, hostname string) error {
	url := fmt.Sprintf("https://%s:%s@domains.google.com/nic/update", username, password)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	queryValues := req.URL.Query()

	queryValues.Set("hostname", hostname)
	queryValues.Set("myip", ip)

	req.URL.RawQuery = queryValues.Encode()

	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}
