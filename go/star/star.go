package star

import (
	"fmt"
	"strings"
)

func PrintTriangle(num int) string {
	var star string = "☆"
	var point string = "○"
	pnt1 := strings.Repeat(point, num)
	fmt.Println(pnt1)
	for i := 1; i <= num; i = i + 2 {
		pnt2 := strings.Repeat(point, (-i+num)*1/2)
		str1 := strings.Repeat(star, i)
		fmt.Println(pnt2, str1, pnt2)
	}
	for j := num - 2; j >= 1; j = j - 2 {
		pnt3 := strings.Repeat(point, (-j+num)*1/2)
		str2 := strings.Repeat(star, j)
		fmt.Println(pnt3, str2, pnt3)
	}
	fmt.Println(pnt1)
	return ""
}

func main() {
	fmt.Println(PrintTriangle(3))
}
