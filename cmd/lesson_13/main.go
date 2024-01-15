package main

import (
	"fmt"

	"github.com/mayerkv/algo/pkg/hash"
)

// example
func main() {
	hashFunc := func(s string) int {
		pearsonHash := hash.PearsonHash(s)
		return int(pearsonHash)
	}
	hashTable := hash.NewHashTable[string, int](hashFunc)
	items := []string{"cat", "dog", "rat", "fox", "bit", "rex"}
	for value, key := range items {
		hashTable.Put(key, value)
	}
	for _, key := range items {
		value, ok := hashTable.Get(key)
		fmt.Println(ok, key, value)
	}
	for i := 0; i < len(items)/2; i++ {
		key := items[i]
		hashTable.Remove(key)
	}
	for _, key := range items {
		value, ok := hashTable.Get(key)
		fmt.Println(ok, key, value)
	}
	fmt.Println(hashTable.IsEmpty(), hashTable.Size())
}
