package towqueue

func Cal(queue1 []int, queue2 []int) {
	top := len(queue1) - 1
	queue1 = queue1[:top]
	data := queue1[top]
	queue2 = append(queue2, data)
}

func Cal2(queue2 []int, queue1 []int) {
	top := len(queue2) - 1
	queue2 = queue2[:top]
	data := queue2[top]
	queue1 = append(queue1, data)
}

func Sum(queue1 []int, queue2 []int) int {
	var sum1 int
	var sum2 int
	var result int
	for i := 0; i < len(queue1); i++ {
		sum1 = sum1 + queue1[i]
	}
	for j := 0; j < len(queue2); j++ {
		sum2 = sum2 + queue2[j]
	}
	for sum1 == sum2 {
		Cal(queue1, queue2)
		Cal2(queue1, queue2)
		result++
	}
	return result
}

func solution(queue1 []int, queue2 []int) int {
	queue1 = []int{3, 2, 7, 2}
	queue2 = []int{4, 6, 5, 1}
	Sum(queue1, queue2)
	return -1
}
