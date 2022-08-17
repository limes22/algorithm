package main

import "fmt"

type ArrayIdx struct {
	Value    int
	Priority int
}

type Queue struct {
	arrayNum []ArrayIdx
	length   int
	front    int
	back     int
}

func (q *Queue) QueueInit(n int) {
	q.length = n
	q.front = 0
	q.back = 0
	q.arrayNum = make([]ArrayIdx, n)
	fmt.Println(q.length)
	fmt.Println("queue init", q.arrayNum)
}

func (q *Queue) EnQueue(inputNumber, priority int) {
	f := q.front
	enq := ArrayIdx{Value: inputNumber, Priority: priority}
	if f == q.length {
		f = 0
	} else if q.arrayNum[f].Value == 0 {
		q.arrayNum[f] = enq
	} else {
		fmt.Println("queue if full")
	}
	q.EnqueueFront()
	q.BubbleSort()
	fmt.Println("enqueue", q.arrayNum, f)
}

func (q *Queue) EnqueueFront() int {
	if q.front == q.length {
		q.front = 0
	} else {
		q.front++
	}
	return q.front
}

func (q *Queue) BubbleSort() {
	f := q.front + 1
	for i := f - 1; i < q.length+(f-1); i++ {
		cur := i
		if i >= q.length {
			cur -= q.length
		}
		for j := i + 1; j < q.length+(i+1); j++ {
			next := j
			if j >= q.length {
				next -= q.length
			} else if q.arrayNum[cur].Priority == q.arrayNum[next].Priority {
				continue
			} else if q.arrayNum[cur].Priority > q.arrayNum[next].Priority {
				tempNum := q.arrayNum[cur]
				q.arrayNum[cur] = q.arrayNum[next]
				q.arrayNum[next] = tempNum
			} else {
				fmt.Println("bubblesort end")
			}
		}
	}
}

func (q *Queue) Dequeue() {
	b := q.back
	zeroq := ArrayIdx{Value: 0, Priority: 0}
	if b == q.length {
		b = 0
	} else if q.arrayNum[b].Value != 0 {
		q.arrayNum[b] = zeroq
	} else {
		fmt.Println("Queue is empty")
	}
	q.DequeueBack()
	fmt.Println("dequeue", q.arrayNum, b)
}

func (q *Queue) DequeueBack() {
	if q.back == q.length {
		q.back = 0
		fmt.Println("queue is empty")
	} else {
		q.back++
	}
}

func main() {
	var q Queue
	q.QueueInit(3)
	q.EnQueue(1, 3)
	q.EnQueue(2, 2)
	q.EnQueue(3, 1)
	q.EnQueue(4, 5)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}
