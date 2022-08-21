package main

import "fmt"

//역전 알고리즘 구현
//회전 시키는 수에 대해 구건을 나누어 reverse로 구현하는 방법
//d = 2 -> 1,2 / 3,4,5,6,7 로 구간을 나누어 reverse 하고 합친다.

//swap을 활용한 reverse 구현
func reverseArr(arr [10]int, start int, end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

//d로 나눠서 역전 알고리즘 수행
func rotateLeft(arr [10]int, d int, n int) {
	reverseArr(arr, 0, d-1)
	reverseArr(arr, d, n-1)
	reverseArr(arr, 0, n-1)
}

func leftRotate(arr []int, d int, n int) {
	i := 0
	temp := arr[i]
	for i = 0; i < gcd(d, n); i++ {
		j := i
		for true {
			k := j + d
			if k >= n {
				k = k - n
			}
			if k == i {
				break
			}
			arr[j] = arr[k]
			j = k
		}
		fmt.Println(j)
		arr[j] = temp
	}
}

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(arr)
	fmt.Println(arr)
	leftRotate(arr, 2, n)
	fmt.Println(arr)
}
