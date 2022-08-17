package queue

import "fmt"

type Queue struct {
	array []int
	front int
	back  int
}

func (q *Queue) QueueInit(n int) []int {
	q.array = make([]int, n, 10)
	q.front = 0
	q.back = 0
	return q.array
}

func (q *Queue) Enqueue(inputNumbers int) []int {
	f := q.front
	if f == len(q.array) {
		f = 0
	} else if q.array[f] == 0 {
		q.array[f] = inputNumbers
	} else {
		fmt.Println("queue is full")
	}
	q.EnqueueFront()
	return q.array
}

func (q *Queue) Dequeue() []int {
	b := q.back
	if b == len(q.array) {
		b = 0
	} else if q.array[b] != 0 {
		q.array[b] = 0
	} else {
		fmt.Println("Queue is empty")
	}
	q.DequeueBack()
	return q.array
}

func (q *Queue) EnqueueFront() int {
	if q.front == len(q.array) {
		q.front = 0
	} else {
		q.front++
	}
	return q.front
}

func (q *Queue) DequeueBack() int {
	if q.back == len(q.array) {
		q.back = 0
		fmt.Println("queue is empty")
	} else {
		q.back++
	}
	return q.back
}

func main() {
	var q Queue
	q.QueueInit(3)
	fmt.Println(q.array)
	fmt.Println(q.Enqueue(1))
	fmt.Println(q.Enqueue(2))
	fmt.Println(q.Enqueue(3))
	fmt.Println(q.Enqueue(4))
	fmt.Println(q.Enqueue(5))
	fmt.Println(q.front)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

}
