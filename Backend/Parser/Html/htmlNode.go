package Html

import (
	"net/url"

	"golang.org/x/net/html"
)

func RewriteHTMLnode(doc *html.Node, base *url.URL) {
	if doc.Type == html.ElementNode {
		for i := range doc.Attr {
			switch doc.Attr[i].Key {
			case "href", "src":
				doc.Attr[i].Val = BrowseURL(doc.Attr[i].Val, base)
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		RewriteHTMLnode(c, base)
	}
}
