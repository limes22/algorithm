package main

import (
	"fmt"
)

type Queue struct {
	array []int
	front int
	back  int
}

func (q *Queue) QueueInit(n int) []int {
	q.array = make([]int, n)
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

func solution(n int, lost []int, reserve []int) int {
	return 0
}

func main() {
	arr1 := [][]int{{1, 4}, {3, 2}, {4, 1}}
	arr2 := [][]int{{3, 3}, {3, 3}}
	fmt.Println(solution(arr1, arr2))
}
