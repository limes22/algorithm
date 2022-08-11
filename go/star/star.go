package star

import (
	"fmt"
	"strings"
)

func main() {
	PrintTriangle(5)
}

func PrintTriangle(num int) string {
	var star string = "*"
	var point string = "."
	for i := 1; i <= num; i += 2 {
		fmt.Printf("%6s"+"%6s\n", strings.Repeat(star, i), strings.Repeat(point, i))
	}
	return ""
}
