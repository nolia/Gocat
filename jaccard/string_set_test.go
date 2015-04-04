package jaccard

import (
	"fmt"
	"log"
	"testing"
)

func Test_NewSetWithValues(t *testing.T) {
	set := NewSetWithValues([]string{"Hello", "world"})

	_, ok := set["Hello"]
	if !ok {
		log.Fatal("'Hello' not found in ", set)
		t.Fail()
	}

	_, ok = set["world"]
	if !ok {
		log.Fatal("'world' not found in ", set)
		t.Fail()
	}
}

func Test_StringSet_Add(t *testing.T) {
	set := NewSet()
	ok := set.Add("hello")

	if !ok {
		t.Fail()
	}
	ok = set.Add("hello")
	if ok {
		t.Fail()
	}
}

func Test_StringSet_Has(t *testing.T) {
	set := NewSet()
	set.Add("music")

	hasMusic := set.Has("music")
	if !hasMusic {
		log.Fatal("'music' not found in ", set)
		t.Fail()
	}
}

func Test_StringSet_Remove(t *testing.T) {
	set := NewSet()
	set.Add("Cat")

	removed := set.Remove("Cat")
	if !removed {
		t.Fail()
	}

	_, stillThere := set["Cat"]
	if stillThere {
		t.Fail()
	}
}

func Test_StringSet_to_String(t *testing.T) {
	set := NewSet()
	set.Add("Hello")
	set.Add("world")
	toString := fmt.Sprintf("%v", set)
	println("String() result: ", toString)

	if toString != "{Hello, world}" {
		log.Fatal("Expected '{Hello, world}', found: ", toString)
		t.Fail()
	}
}
