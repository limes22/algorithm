package main

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

func main() {

	rootTree := InitTree()

	rootTree.AddNode("", "first parent Node", "data1")
	rootTree.AddNode("", "second parent Node", "data2")

	rootTree.AddNode("first parent Node", "first parent child Node", "data3")

}
