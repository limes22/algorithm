package Practice

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	before := 0
	answer := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Sort(sort.IntSlice(arr))
	for i := 0; i < n; i++ {
		answer += before + arr[i]
		before += arr[i]
	}
	fmt.Println(answer)
}
