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
