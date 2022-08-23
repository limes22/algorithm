package revers_array

import "fmt"

//지정된 배열에서 하나씩 회전을 해서 i*arr[i]의 합이 가장 컸을 때 값을 출력하는 문제

func MaxVal(arr []int, n int) int {
	arrSum := 0 //arr[i]의 전체값
	curSum := 0 //i*arr[i]의 전체값
	for i := 0; i < n; i++ {
		arrSum = arrSum + arr[i]
		curSum = curSum + (i * arr[i])
	}
	maxSum := curSum

	for j := 1; j < n; j++ {
		curSum = curSum + arrSum - n*arr[n-j]
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
	arr := []int{1, 20, 2, 10}
	n := len(arr)
	fmt.Println(MaxVal(arr, n))
}
