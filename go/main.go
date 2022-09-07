package main

import (
	"fmt"
)

func main() {
	var n int
	coin := 0
	fmt.Scan(&n)
	if n >= 50000 {
		coin += n / 50000
		n = n % 50000
	}
	if n >= 10000 {
		coin += n / 10000
		n = n % 10000
	}
	if n >= 5000 {
		coin += n / 5000
		n = n % 5000
	}
	if n >= 1000 {
		coin += n / 1000
		n = n % 1000
	}
	if n >= 500 {
		coin += n / 500
		n = n % 500
	}
	if n >= 100 {
		coin += n / 100
		n = n % 100
	}
	if n >= 50 {
		coin += n / 50
		n = n % 50
	}
	if n >= 10 {
		coin += n / 10
		n = n % 10
	}
	fmt.Println(coin, n)
}
