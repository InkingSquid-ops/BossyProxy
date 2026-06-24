package Html

import (
	"bytes"
	"io"
	"net/url"

	"golang.org/x/net/html"
)

func rewriteHTML(r io.Reader, base *url.URL) ([]byte, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	rewriteHTMLnode(doc, base)

	var buf bytes.Buffer

	if err := html.Render(&buf, doc); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
