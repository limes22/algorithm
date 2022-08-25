package tree

type Tree struct {
	root []*Node
}

type Node struct {
	parents  []*Node
	child    []*Node
	nodeName string
	data     string
}

var tree Tree

func InitTree() *Tree {
	tree = Tree{}
	return &tree
}

func (t *Tree) AddNode(parentNodeName string, nodeName string, data string) {
	node := Node{
		parents:  []*Node{},
		child:    []*Node{},
		nodeName: nodeName,
		data:     data,
	}

	for _, parentNodeAddr := range tree.root {
		if (*parentNodeAddr).nodeName == parentNodeName {
			node.parents = append(node.parents, parentNodeAddr)
			(*parentNodeAddr).child = append((*parentNodeAddr).child, &node)
			return
		}
	}

	tree.root = append(tree.root, &node)
}

//func (t *Tree) ShowNode () {
//	var n Node
//	fmt.Println("current nodeName:", n.nodeName)
//	if n.child == nil {
//		return;
//	} else {
//		fmt.Println("childs:", n.child)
//		for _, NodeAddr := range tree.root {
//			t.ShowNode("child:", NodeAddr)
//		}
//	}
//}

func main() {

	rootTree := InitTree()

	rootTree.AddNode("", "first parent Node", "data1")
	rootTree.AddNode("", "second parent Node", "data2")

	rootTree.AddNode("first parent Node", "first parent child Node", "data3")
	rootTree.AddNode("first parent Node", "second parent child Node", "data4")
	rootTree.AddNode("first parent child Node", "first parent child  child Node", "data5")

}
