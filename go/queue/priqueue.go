package priqueue

import "fmt"

type Array []interface{}

type Queue struct {
	ArrayNum []int
	ArrayPri []int
	Length   int
	Front    int
	Back     int
}

func (q *Queue) LengthInit(n int) int {
	q.Length = n
	fmt.Println("length", q.Length)
	q.QueueNumInit()
	q.QueuePriInit()
	fmt.Println(q.ArrayNum, q.ArrayPri)
	return q.Length
}

func (q *Queue) QueueNumInit() []int {
	q.ArrayNum = make([]int, q.Length)
	return q.ArrayNum
}

func (q *Queue) QueuePriInit() []int {
	q.ArrayPri = make([]int, q.Length)
	return q.ArrayPri
}

func (q *Queue) Enqueue(inputNumbers int, priority int) []int {
	f := q.Front
	if f == q.Length {
		f = 0
		q.ArrayNum[f] = inputNumbers
		q.ArrayPri[f] = priority
	} else if q.ArrayNum[f] == 0 {
		q.ArrayNum[f] = inputNumbers
		q.ArrayPri[f] = priority
	} else {
		fmt.Println("queue is full")
	}
	q.EnqueueFront()
	q.BubbleSort()
	fmt.Println("enequqe", q.ArrayNum, q.ArrayPri)
	q.ReturnPri()
	return q.ArrayNum
}

func (q *Queue) ReturnPri() []int {
	return q.ArrayPri
}

func (q *Queue) Dequeue() []int {
	b := q.Back
	if b == q.Length {
		b = 0
	} else if q.ArrayNum[b] != 0 {
		q.ArrayNum[b] = 0
		q.ArrayPri[b] = 0
	} else {
		fmt.Println("Queue is empty")
	}
	q.DequeueBack()
	fmt.Println("dequqe", q.ArrayNum, q.ArrayPri)
	q.ReturnPri()
	return q.ArrayNum
}

func (q *Queue) EnqueueFront() int {
	if q.Front == q.Length {
		q.Front = 0
	} else {
		q.Front++
	}
	return q.Front
}

func (q *Queue) DequeueBack() int {
	if q.Back == q.Length {
		q.Back = 0
		fmt.Println("queue is empty")
	} else {
		q.Back++
	}
	return q.Back
}

func (q *Queue) BubbleSort() []int {
	f := q.Front
	for i := f - 1; i < q.Length+(f-1); i++ {
		cur := i
		if i >= q.Length {
			cur -= q.Length
		}
		for j := i + 1; j < q.Length+(i+1); j++ {
			next := j
			if j >= q.Length {
				next -= q.Length
			} else if q.ArrayPri[cur] == 0 {
				continue
			} else if q.ArrayPri[cur] > q.ArrayPri[next] {
				tempNum := q.ArrayNum[cur]
				tempPri := q.ArrayPri[cur]
				q.ArrayNum[cur] = q.ArrayNum[next]
				q.ArrayPri[cur] = q.ArrayPri[next]
				q.ArrayNum[next] = tempNum
				q.ArrayPri[next] = tempPri
			} else {
				fmt.Println("bubblesort end")
			}
		}
	}
	q.ReturnPri()
	return q.ArrayNum
}

func main() {
	var q Queue
	q.LengthInit(3)
	q.Enqueue(1, 3)
	q.Enqueue(2, 4)
	fmt.Println(q.Front)
	q.Enqueue(3, 5)
}
