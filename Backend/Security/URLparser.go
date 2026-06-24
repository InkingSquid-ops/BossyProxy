package security

import (
	"errors"
	"net/url"
)

func ValidateURL(raw string) error {
	u, err := url.Parse(raw)

	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("Invaild scheme")
	}

	return nil
}
