package Html

import (
	"net/url"

	"golang.org/x/net/html"
)

func rewriteHTMLnode(doc *html.Node, base *url.URL) error {
	if doc.Type == html.ElementNode {
		for i := range doc.Attr {
			switch doc.Attr[i].Key {
			case "href", "src":
				doc.Attr[i].Key = BrowseURL(doc.Attr[i].Key, base)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		rewriteHTMLnode(doc, base)
	}
}
