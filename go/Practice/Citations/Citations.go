package Citations

import (
	"fmt"
	"sort"
)

func solution(citations []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	for h := 0; h < len(citations); h++ {
		// fmt.Println(citations[h], "번 이상 인용된 논문이 ", h+1, "개 있다.")

		if citations[h] < h+1 {
			return h
		}
	}

	return len(citations)
}

func Citations() {
	arr := []int{3, 0, 6, 1, 5}
	fmt.Println(solution(arr))
}
