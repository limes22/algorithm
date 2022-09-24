package Tournament

import (
	"fmt"
)

func solution(n int, a int, b int) int {
	answer := 0
	a--
	b--

	for {
		if a == b {
			break
		}
		answer += 1
		a /= 2
		b /= 2
	}

	return answer
}

func Tournament() {
	fmt.Println(solution(8, 4, 7))
}
