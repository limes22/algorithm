package Bracket

import (
	"strings"
)

func solution(s string) int {
	answer := 0
	for i := 0; i < len(s); i++ {
		if isRight(s) {
			answer++
		}
		s = rotate(s)
	}

	return answer
}

func rotate(s string) string {
	head := s[0]
	rotated := s[1:]
	return string(rotated) + string(head)
}
func isRight(s string) bool {
	stack := []string{}
	open := "[{("
	close := "]})"
	L := len(s)
	for i := 0; i < L; i++ {
		if strings.Contains(open, string(s[i])) {
			stack = append(stack, string(s[i]))
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if strings.Index(open, stack[len(stack)-1]) != strings.Index(close, string(s[i])) {
			return false
		}
		stack = stack[0 : len(stack)-1]
	}
	return len(stack) == 0
}
