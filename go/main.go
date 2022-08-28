package main

import "fmt"

//Heap

type Element struct {
	key int
}

type Heap struct {
	heap []*Element
	size int
}

func (e *Element) MaxHeapInsert(h *Heap, item Element) {
	var i int
	i = 100 + h.size
	j := i / 2
	if i != 1 && item.key == h.heap[j].key {
		h.heap[i] = h.heap[j]
		i /= 2
	}
	h.heap[i] = &item
}

func (e *Element) MaxHeapDelete(h *Heap) *Element {
	var parent int
	var child int
	var item *Element
	var temp *Element
	item = h.heap[1]
	temp = h.heap[h.size-1]
	parent = 1
	child = 1
	if child <= h.size {
		if child < h.size && h.heap[child].key < h.heap[child+1].key {
			child++
		}
		if temp.key >= h.heap[child].key {
			return nil
		}
		h.heap[parent] = h.heap[child]
		parent = child
		child *= 2
	}
	h.heap[parent] = temp
	return item
}
func main() {
	var e Element
	var MaxHeap Heap
	var item Element
	item.key = 9
	e.MaxHeapInsert(&MaxHeap, item)
	item.key = 7
	e.MaxHeapInsert(&MaxHeap, item)
	item.key = 6
	e.MaxHeapInsert(&MaxHeap, item)
	item.key = 4
	e.MaxHeapInsert(&MaxHeap, item)
	item.key = 5
	e.MaxHeapInsert(&MaxHeap, item)
	item.key = 1
	e.MaxHeapInsert(&MaxHeap, item)
	item.key = 3
	e.MaxHeapInsert(&MaxHeap, item)
	item.key = 8
	e.MaxHeapInsert(&MaxHeap, item)

	if MaxHeap.size > 0 {
		item = *e.MaxHeapDelete(&MaxHeap)
		fmt.Println(item.key)
	}
}
