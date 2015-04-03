package crawl

import (
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func prepare(s string, t *testing.T) *html.Node {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	return doc
}

func TestCollectLinks(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc := prepare(s, t)
	collected := CollectLinks(doc)

	if len(collected) != 2 {
		log.Fatal("Expected 2 links, found ", len(collected))
		t.Fail()
	}

	if collected[0] != "foo" {
		log.Fatal("Expected foo, but found ", collected[0])
	}

	if collected[1] != "/bar/baz" {
		log.Fatal("Expected bar/baz, but found ", collected[1])
	}
}

func TestCollectText(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc := prepare(s, t)
	collectedText := CollectText(doc)

	log.Println("result:\n", collectedText)

	if !strings.Contains(collectedText, "Foo") {
		log.Fatal("Foo not found")
		t.Fail()
	}

	if !strings.Contains(collectedText, "BarBaz") {
		log.Fatal("Foo not found")
		t.Fail()
	}
}
