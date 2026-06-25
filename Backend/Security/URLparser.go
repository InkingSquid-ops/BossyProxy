package Security

import (
	"strings"
)

func ValidateURL(raw string) string {
	if !strings.Contains(raw, "https://") && !strings.Contains(raw, "http://") {
		newRaw := "https://" + raw
		return newRaw
	}

	return raw
}
