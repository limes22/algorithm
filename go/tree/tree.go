package tree

type Node struct {
	parents  []*Node
	child    []*Node
	nodeName string
}

type NodeRoot struct {
	rootNode []*Node
}

func (n *NodeRoot) Tree(nodeName string) *NodeRoot {
	if n.rootNode == nil {
		n.rootNode = &Node{nodeName: nodeName}
	}
}
