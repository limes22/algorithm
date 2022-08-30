package Trie

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	r    = bufio.NewReader(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	t, n int
)

func main() {
	defer w.Flush()
	fmt.Fscan(r, &t)
	for t > 0 {
		t--
		fmt.Fscan(r, &n)
		arr := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &arr[i])
		}

		// 문자열 오름차순으로 정렬
		sort.Strings(arr)

		check := true
	run:
		// 접두어인지 판별
		for i := n - 1; i > 0; i-- {
			if strings.HasPrefix(arr[i], arr[i-1]) {
				check = false
				break run
			}
		}

		if check {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
