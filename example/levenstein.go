package main

import (
	"fmt"
	levenshtein "github.com/creasty/go-levenshtein"
	"bitbucket.org/egorsteam/bk-tree/internal"
)

type word string

// Distance calculates hamming distance.
func (x word) Distance(e internal.ObjectTree) int {
	a := string(x)
	b := string(e.(word))

	return levenshtein.Distance(a, b)
}

func main() {
	var tree internal.Tree
	// add words
	words := []string{"apple", "banana", "orange", "peach", "bean", "tomato", "egg", "pineapple"}
	for _, w := range words {
		tree.Insert(word(w))
	}

	// spell check
	results := tree.Search(word("peacn"), 2)
	fmt.Println("Input is peacn. Did you mean:")
	for _, result := range results {
		fmt.Printf("\t%s (distance: %d)\n", result.Object.(word), result.Distance)
	}
}