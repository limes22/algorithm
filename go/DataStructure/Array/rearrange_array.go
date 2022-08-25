package revers_array

import "fmt"

func fix(A []int, len int) {
	for i := 0; i < len; i++ {
		if A[i] != -1 && A[i] != i {
			x := A[i]
			if A[x] != -1 && A[x] != x {
				y := A[x]
				x = y
			}
			A[x] = x
			if A[i] != i {
				A[i] = -1
			}
		}
	}
}

func main() {
	A := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	len := len(A)
	fix(A, len)
	fmt.Println(A)
}
