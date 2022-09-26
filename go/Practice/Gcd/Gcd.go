package Gcd

import "fmt"

func gcd(x int, y int) int {
	if x%y == 0 {
		return y
	} else {
		return gcd(y, x%y)
	}
}

func lcm(x int, y int) int {
	return (x * y) / gcd(x, y)
}

func solution(arr []int) int {
	answer := 1
	for i := 1; i < len(arr); i++ {
		answer = lcm(answer, lcm(arr[i-1], arr[i]))
	}

	return answer
}

func Gcd() {
	arr := []int{2, 6, 8, 14}
	fmt.Println(solution(arr))
}
