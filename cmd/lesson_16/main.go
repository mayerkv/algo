package main

import (
	"log"

	"github.com/mayerkv/algo/pkg/trie"
)

func main() {
	t := trie.Constructor()
	t.Insert("apple")
	assertIsTrue(t.Search("apple"))
	assertIsFalse(t.Search("app"))
	assertIsTrue(t.StartsWith("app"))
	t.Insert("app")
	assertIsTrue(t.Search("app"))

	log.Println("Test success")
}

func assertIsTrue(value bool) {
	if value != true {
		log.Fatalf("Value must be true, %v given", value)
	}
}

func assertIsFalse(value bool) {
	if value != false {
		log.Fatalf("Value must be false, %v given", value)
	}
}
