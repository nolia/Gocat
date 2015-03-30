package gocat

import (
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCollectLinks(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
		t.Fail()
	}
	collected := CollectLinks(doc)

	if len(collected) != 2 {
		log.Fatal("Expected 2 links, found", len(collected))
		t.Fail()
	}

	if collected[0] != "foo" {
		log.Fatal("Expected foo, but found ", collected[0])
	}

	if collected[1] != "/bar/baz" {
		log.Fatal("Expected bar/baz, but found ", collected[1])
	}
}
