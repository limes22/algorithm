package JadenCase

import "strings"

func solution(s string) string {
	var answer string
	t := true
	for _, i := range s {
		if i == ' ' {
			t = true
			answer += " "
		} else if t {
			answer += strings.ToUpper(string(i))
			t = false
		} else {
			answer += strings.ToLower(string(i))
		}
	}
	return answer
}
