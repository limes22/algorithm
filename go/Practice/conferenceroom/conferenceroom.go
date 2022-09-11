package conferenceroom

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	arr := make([][2]int, n)
	answer := 0
	endTime := 0

	// input 받는 과정
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &arr[i][0], &arr[i][1])
	}

	// 끝나는 시간으로 정렬 (같을 시 시작하는 시간 순으로)
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})

	// 회의실을 사용가능한지 확인하는 로직
	for i := 0; i < n; i++ {
		if endTime <= arr[i][0] && endTime <= arr[i][1] {
			answer++
			endTime = arr[i][1]
		}
	}
	fmt.Println(answer)
}
