package Html

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

func rewriteHTML(r io.Reader, base *url.URL) ([]byte, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	rewriteHTMLnode(doc)

}
