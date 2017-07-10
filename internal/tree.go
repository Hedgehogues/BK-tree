package internal

// Тип вынесен в интерфейс, так как в ряде случаев может быть необходим целочисленный тип. В иных случаях --
// вещественный
type TypeOfDistance interface {
	More(TypeOfDistance) bool
	Plus(TypeOfDistance) TypeOfDistance
	Minus(TypeOfDistance) TypeOfDistance
}

type ObjectTree interface {
	Distance(ObjectTree) TypeOfDistance
}

type Node struct {
	objectTree ObjectTree
	children   map[TypeOfDistance]*Node
}

type Tree struct {
	Root *Node
}

type Result struct {
	Distance TypeOfDistance
	Object   ObjectTree
}

func (node *Node) insert(object ObjectTree) {
	dist := node.objectTree.Distance(object)
	for newNode, ok := node.children[dist]; ok; newNode, ok = node.children[dist] {
		node = newNode
		dist = node.objectTree.Distance(object)
	}

	node.children[dist] = &Node{
		objectTree: object,
		children:   map[TypeOfDistance]*Node{},
	}
	return
}

func (tree *Tree) Insert(object ObjectTree) {
	if tree.Root == nil {
		tree.Root = &Node{
			children:   map[TypeOfDistance]*Node{},
			objectTree: object,
		}
		return
	}
	tree.Root.insert(object)
}

func (tree *Tree) Search(object ObjectTree, tolerance TypeOfDistance) []Result {
	results := make([]Result, 0)
	if tree.Root == nil {
		return results
	}
	candidates := []*Node{tree.Root}
	for len(candidates) != 0 {
		lastElement := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		dist := lastElement.objectTree.Distance(object)
		if !dist.More(tolerance) {
			// dist <= tolerance
			results = append(results, Result{
				Distance: dist,
				Object:   lastElement.objectTree,
			})
		}

		low, high := dist.Minus(tolerance), dist.Plus(tolerance)
		for distance, candidate := range lastElement.children {
			if !low.More(distance) && !distance.More(high) {
				// low <= distance && distance <= high
				candidates = append(candidates, candidate)
			}
		}
	}
	return results
}
