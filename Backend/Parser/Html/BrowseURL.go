package Html

import (
	"net/url"
)

func BrowseURL(TargetURL string, base *url.URL) string {
	u, err := url.Parse(TargetURL)

	if err != nil {
		return TargetURL
	}

	newBase := base.ResolveReference(u)

	return "/?url=" + url.QueryEscape(newBase.String())
}
