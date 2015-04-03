package jaccard

import (
	"fmt"
	"log"
	"testing"
)

func TestSplitWord(t *testing.T) {
	s := "Hello, world!"
	words := SplitWords(s)
	fmt.Println("Split result:\n", words)

	if len(words) != 2 {
		log.Fatal("Expeced len = 2, found len = ", len(words))
		t.Fail()
	}
	if words[0] != "Hello" {
		log.Fatal("Expected 'Hello', found ", words[0])
		t.Fail()
	}
	if words[1] != "world" {
		log.Fatal("Expected 'world', found ", words[1])
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
