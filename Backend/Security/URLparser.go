package Security

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func ValidateURL(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)

	if !strings.HasPrefix(trimmed, "http://") && !strings.HasPrefix(trimmed, "https://") {
		testHTTPS := "https://" + trimmed

		client := &http.Client{
			Timeout: 4 * time.Second,
		}

		_, err := client.Head(testHTTPS)
		if err != nil {
			// HTTPS failed, fallback to HTTP
			trimmed = "http://" + trimmed
		} else {
			// HTTPS succeeded
			trimmed = testHTTPS
		}
	}

	u, err := url.Parse(trimmed)
	if err != nil {
		return "", err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return "", errors.New("Invalid scheme")
	}

	return trimmed, nil
}