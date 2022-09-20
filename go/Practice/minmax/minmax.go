package minmax

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solution(s string) string {
	strs := strings.Fields(s)

	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	sort.Ints(ints)

	return fmt.Sprintf("%d %d", ints[0], ints[len(ints)-1])
}
