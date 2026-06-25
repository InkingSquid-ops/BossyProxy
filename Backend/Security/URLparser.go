package Security

import (
	"net/url"
	"strings"
)

func ValidateURL(raw string) string {
	u, err := url.Parse(raw)

	if err != nil {
		return ""
	}

	switch strings.ToLower(u.Scheme) {
	case "":
		return "https://" + raw

	case "file", "ftp":
		return ""

	case "javascript", "mailto", "tel":
		return ""

	default:
		return raw
	}
}

