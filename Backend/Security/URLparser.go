package Security

import (
	"fmt"
	"strings"
)

func ValidateURL(raw string) string {
	switch {
	case strings.Contains(raw, "file://"):
		return fmt.Sprintf("Invalid url: %s", raw)

	case strings.Contains(raw, "ftp://"):
		return fmt.Sprintf("Invalid url: %s", raw)

	default:
		if !strings.Contains(raw, "https://") && !strings.Contains(raw, "http://") {
			newRaw := "https://" + raw
			return newRaw
		}
	}

	return raw
}
