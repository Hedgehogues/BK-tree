package bk_tree

type ObjectVertex interface {
	Distance(ObjectVertex) int
}

type Node struct {
	objectVertex ObjectVertex
	children     map[int]*Node
}

type Tree struct {
	root *Node
}

type Result struct {
	Distance int
	Object   ObjectVertex
}

func (node *Node) insert(object ObjectVertex) {
	dist := node.objectVertex.Distance(object)
	for newNode, ok := node.children[dist]; ok; newNode, ok = node.children[dist] {
		node = newNode
		dist = node.objectVertex.Distance(object)
	}
	node.children[dist] = &Node{
		objectVertex: object,
		children:     map[int]*Node{},
	}
	return
}

func (tree *Tree) Insert(object ObjectVertex) {
	if tree.root == nil {
		tree.root = &Node{
			children:     map[int]*Node{},
			objectVertex: object,
		}
		return
	}
	tree.root.insert(object)
}

func (tree *Tree) Search(object ObjectVertex, tolerance int) []Result {
	results := make([]Result, 0)
	if tree.root == nil {
		return results
	}
	candidates := []*Node{tree.root}
	for len(candidates) != 0 {
		candidates = candidates[:len(candidates)-1]
		d := candidates[len(candidates)-1].objectVertex.Distance(object)
		if d <= tolerance {
			results = append(results, Result{
				Distance: d,
				Object:   candidates[len(candidates)-1].objectVertex,
			})
		}

		low, high := d-tolerance, d+tolerance
		for distance, c := range candidates[len(candidates)-1].children {
			if low <= distance && distance <= high {
				candidates = append(candidates, c)
			}
		}
	}
	return results
}
