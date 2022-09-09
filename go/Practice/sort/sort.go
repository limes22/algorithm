package sorting

import (
	"fmt"
	"sort"
)

func sorting() {

	answer := 0
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Sort(sort.IntSlice(a))
	before := a[0]
	for i := 1; i < n; i++ {
		answer += before + a[i]
		before += a[i]
	}
	fmt.Println(answer)
	fmt.Println(a)
}
