package main

import "fmt"

//stack
//LIFO(Last In First Out, 후입선출): 가장 나중에 들어온 것이 가장 먼저 나옴
//함수의 콜스택, 문자열 역순 출력, 연ㅅ안자 후위 표기법
//데이터 넣음 : push()
//데이터 최상위 값 뺌 : pop()
//비어있는지 확인 : isEmpty()
//꽉차있는지 확인 : isFull()

type Node struct {
	data int
	next *Node
}

type Main struct {
	head *Node
}

func (n *Node) NodeInit(d int) *Node {
	node := Node{
		data: d,
		next: nil,
	}
	return &node
}

func (m *Main) Push(new_data int) {
	new_node := new(Node)
	new_node.data = new_data
	new_node.next = m.head
	m.head = new_node
}

func (m *Main) InsertAfter(prev_node *Node, new_data int) {
	if prev_node == nil {
		fmt.Println("The given Previous node canot be null")
		return
	}
	new_node := new(Node)
	new_node.data = new_data
	new_node.next = prev_node.next
	prev_node.next = new_node
}

func (m *Main) Append(new_data int) {
	new_node := new(Node)
	new_node.data = new_data
	if m.head == nil {
		m.head = new(Node)
		m.head.data = new_data
		return
	}
	new_node.next = nil
	last := m.head
	if last.next != nil {
		last = last.next
		last.next = new_node
		return
	}
}

func (m *Main) PrintList() {
	tnode := m.head
	if tnode != nil {
		fmt.Println(tnode.data, "")
		tnode = tnode.next
	}
}
func main() {
	var m Main
	m.Append(8)
	m.Push(7)
	m.Push(1)
	m.Append(4)
	m.InsertAfter(m.head.next, 8)
	m.PrintList()
}
