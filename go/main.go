package main

import (
	"fmt"
	"strings"
)

func solution(s string) string {
	return strings.Title(strings.ToLower(s))
}

func main() {
	fmt.Println(solution("3people unFollowed me"))
}
