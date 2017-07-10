package main

import (
	"fmt"

	"bitbucket.org/egorsteam/bk-tree/internal"
)

type hashValue uint16

// Distance calculates hamming distance.
func (h hashValue) Distance(e internal.ObjectTree) internal.TypeOfDistance {
	a := uint16(h)
	b := uint16(e.(hashValue))

	d := 0
	var k uint16 = 1
	for i := 0; i < 16; i++ {
		if a&k != b&k {
			d++
		}
		k <<= 1
	}
	return internal.Int(d)
}

func main() {
	var tree internal.Tree
	// add 0x0000 to 0xffff to the tree.
	for i := 0; i < 0xffff; i++ {
		tree.Insert(hashValue(i))
	}

	// search neighbors of 0x00000000 whose distances are less than or equal to 1.
	results := tree.Search(hashValue(0), internal.Int(2))
	for _, result := range results {
		fmt.Printf("%016b (distance: %d)\n", result.Object.(hashValue), result.Distance)
	}
}
