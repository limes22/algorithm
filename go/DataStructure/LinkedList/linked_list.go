package LinkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type Main struct {
	head *Node
}

func (n *Node) NodeInit(d int) Node {
	node := Node{
		data: d,
		next: nil,
	}
	return node
}

func (n *Node) PrintList() {
	var m Main
	n = m.head
	if n != nil {
		fmt.Println(n.data, "")
		n = n.next
	}
}

func main() {
	var n Node
	llist := new(Main)
	first := n.NodeInit(1)
	llist.head = &first
	second := n.NodeInit(2)
	third := n.NodeInit(3)

	llist.head.next = &second
	second.next = &third
	fmt.Println(llist)
	n.PrintList()
}
