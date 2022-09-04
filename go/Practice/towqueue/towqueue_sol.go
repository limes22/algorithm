package towqueue

import "fmt"

func sumOfQueue(queue []int) int {
	sum := 0
	for _, num := range queue {
		sum += num
	}
	return sum
}

func solution(queue1 []int, queue2 []int) int {

	count := 0
	for !(sumOfQueue(queue1) == sumOfQueue(queue2)) {
		if sumOfQueue(queue1) > sumOfQueue(queue2) {
			queue2 = append(queue2, queue1[0])
			queue1 = queue1[1:len(queue1)]
		} else {
			queue1 = append(queue1, queue2[0])
			queue2 = queue2[1:len(queue2)]
		}
		count += 1
		if len(queue1) == 0 || len(queue2) == 0 {
			return -1
		}
	}
	return count
}

func main() {
	queue1 := []int{3, 2, 7, 2}
	queue2 := []int{4, 6, 5, 1}
	fmt.Println(solution(queue1, queue2))
}
