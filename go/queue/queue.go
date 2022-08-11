package queue

import "fmt"

type Queue struct {
	Array []int
	Front int
	Back  int
}

func (q *Queue) QueueInit(n int) []int {
	q.Array = make([]int, n, 10)
	q.Front = 0
	q.Back = 0
	return q.Array
}

func (q *Queue) Enqueue(inputNumbers int) []int {
	f := q.Front
	if f == len(q.Array) {
		f = 0
	} else if q.Array[f] == 0 {
		q.Array[f] = inputNumbers
	} else {
		fmt.Println("queue is full")
	}
	q.EnqueueFront()
	return q.Array
}

func (q *Queue) Dequeue() []int {
	b := q.Back
	if b == len(q.Array) {
		b = 0
	} else if q.Array[b] != 0 {
		q.Array[b] = 0
	} else {
		fmt.Println("Queue is empty")
	}
	q.DequeueBack()
	return q.Array
}

func (q *Queue) EnqueueFront() int {
	if q.Front == len(q.Array) {
		q.Front = 0
	} else {
		q.Front++
	}
	return q.Front
}

func (q *Queue) DequeueBack() int {
	if q.Back == len(q.Array) {
		q.Back = 0
		fmt.Println("queue is empty")
	} else {
		q.Back++
	}
	return q.Back
}

func main() {
	var q Queue
	q.QueueInit(3)
	fmt.Println(q.Array)
	fmt.Println(q.Enqueue(1))
	fmt.Println(q.Enqueue(2))
	fmt.Println(q.Enqueue(3))
	fmt.Println(q.Enqueue(4))
	fmt.Println(q.Enqueue(5))
	fmt.Println(q.Front)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

}
