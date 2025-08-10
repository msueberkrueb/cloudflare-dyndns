package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func GetCurrentIp() (string, error) {
	req, err := http.Get("https://1.1.1.1/cdn-cgi/trace")

	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	for element := range strings.SplitSeq(string(body[:]), "\n") {
		parts := strings.SplitN(element, "=", 2)

		if len(parts) != 2 {
			return "", errors.New("IPError: CloudFlare returned invalid body")
		}

		key, value := parts[0], parts[1]

		if key != "ip" {
			continue
		}

		return value, nil
	}

	return "", errors.New("IPError: CloudFlare did not provide an ip address")
}
