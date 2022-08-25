package LinkedList

import (
	"fmt"
)

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

func (n *Node) Push(new_data int) *Node {
	var m Main
	n.data = new_data
	new_node := new(Node)
	new_node.next = m.head
	m.head = new_node
	return m.head
}

func (n *Node) InsertAfter(prev_node *Node, new_data int) *Node {
	if prev_node == nil {
		fmt.Println("The given Previous node canot be null")
		return nil
	}
	n.data = new_data
	new_node := new(Node)
	new_node.next = prev_node.next
	prev_node.next = new_node
	return prev_node.next
}

func (n *Node) Append(new_data int) *Node {
	var m Main
	n.data = new_data
	new_node := new(Node)
	if m.head == nil {
		m.head = new(Node)
		return nil
	}
	new_node.next = nil
	last := m.head
	if last.next != nil {
		last = last.next
	}
	last.next = new_node
	return last.next
}

func (n *Node) PrintList() {
	var m Main
	tnode := m.head
	if tnode != nil {
		fmt.Println(tnode.data, "")
		tnode = tnode.next
	}
}
func main() {
	var n Node
	llist := new(Main)
	llist.head = n.Append(6)
	llist.head = n.Push(7)
	llist.head = n.Push(1)
	llist.head = n.Append(4)
	llist.head = n.InsertAfter(llist.head.next, 8)
	fmt.Println("Created Linked list is")
	n.PrintList()
}
