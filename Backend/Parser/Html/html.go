package Html

import (
	"bytes"
	"io"
	"net/url"

	"golang.org/x/net/html"
)

func RewriteHTML(r io.Reader, base string) ([]byte, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	result, err := url.Parse(base)

	if err != nil {
		return nil, err
	}

	RewriteHTMLnode(doc, result)

	var buf bytes.Buffer

	if err := html.Render(&buf, doc); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

