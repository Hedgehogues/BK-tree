package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type wordType string

func (x wordType) Distance(e ObjectTree) TypeOfDistance {
	return Int(0)
}

func TestTree_InsertSearchDistance(t *testing.T) {
	var tree Tree
	words := []string{"apple", "banana", "orange", "peach", "bean", "tomato", "egg", "pineapple"}
	for _, w := range words {
		tree.Insert(wordType(w))
	}

	answer := map[Result]struct{}{}
	for _, word := range words {
		answer[Result{
			Distance: Int(0),
			Object:   wordType(word),
		}] = struct{}{}
	}

	results := tree.Search(wordType("TEST"), Int(1))
	for _, word := range results {
		_, ok := answer[word]
		assert.EqualValues(t, true, ok)
	}
}
