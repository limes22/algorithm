package main

import (
	"fmt"
)

func main() {
	var s, n int
	sum := 0
	fmt.Scan(&s)

	for i := 1; true; i++ {
		sum += i
		if sum > s {
			n = i - 1
			break
		}
	}
	fmt.Println(n)
}
