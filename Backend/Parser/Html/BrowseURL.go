package Html

import (
	"net/url"
)

func BrowseURL(targetURL string, base *url.URL) string {
	if base == nil {
		return targetURL
	}

	u, err := url.Parse(targetURL)
	if err != nil {
		return targetURL
	}

	resolved := base.ResolveReference(u)

	return "/browse/?url=" +
		url.QueryEscape(resolved.String())
}
