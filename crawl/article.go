package crawl

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

func dsfWith(doc *html.Node, v func(*html.Node)) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		v(n)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func CollectLinks(doc *html.Node) []string {
	var links []string = make([]string, 0)

	dsfWith(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					links = append(links, a.Val)
					break
				}
			}
		}
	})

	return links
}

func CollectText(doc *html.Node) string {
	var buf bytes.Buffer
	dsfWith(doc, func(n *html.Node) {
		if n.Type == html.TextNode {
			// Add next line
			if buf.Len() > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(n.Data)
		}
	})

	return buf.String()
}
