package revers_array

import "fmt"

//기본 회전 알고리즘 구현

//왼쪽으로 한번 회전
func leftRotatebyOne(arr []int, n int) {
	i := 0
	temp := arr[i]
	for i = 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[i] = temp
}

func rotateLeft(arr []int, d int, n int) {
	for i := 0; i < d; i++ {
		leftRotatebyOne(arr, n)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(arr)
	fmt.Println(arr)
	rotateLeft(arr, 2, n)
	fmt.Println(arr)
}
