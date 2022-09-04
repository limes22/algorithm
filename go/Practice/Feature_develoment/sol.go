package Feature_develoment

func solution(progresses []int, speeds []int) []int {

	var max, count int = 0, 0
	result := []int{}

	for idx, completTime := range progresses {
		var remainPeriod int = 100 - completTime
		var completPeriod int = (remainPeriod / speeds[idx])

		if remainPeriod%speeds[idx] != 0 {
			completPeriod++
		}

		if max < completPeriod {
			result = append(result, count)
			count = 1
			max = completPeriod
		} else {
			count++
		}
	}

	result = append(result[:0], result[1:]...)
	return append(result, count)
}
