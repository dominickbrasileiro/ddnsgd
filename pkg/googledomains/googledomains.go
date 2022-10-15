package googledomains

import (
	"fmt"
	"io"
	"net/http"
)

type ApiError struct {
	Code string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("could not update ddns (%s)", e.Code)
}

func UpdateDDNS(ip string, username string, password string, hostname string) error {
	url := fmt.Sprintf("https://%s:%s@domains.google.com/nic/update", username, password)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return err
	}

	queryValues := req.URL.Query()

	queryValues.Set("hostname", hostname)
	queryValues.Set("myip", ip)

	req.URL.RawQuery = queryValues.Encode()

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	errorCodes := []string{
		"nohost",
		"badauth",
		"notfqdn",
		"badagent",
		"abuse",
		"911",
		"conflict A",
		"conflict AAAA",
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	bodyContent := string(bodyBytes)

	for _, code := range errorCodes {
		if bodyContent == code {
			return ApiError{code}
		}
	}

	return nil
}
