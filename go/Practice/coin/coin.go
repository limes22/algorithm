package coin

import "fmt"

func coin() {
	var n, k, coin, now int
	coins := make([]int, n)
	fmt.Scanf("%d %d", &n, &k)
	answer := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&coin)
		coins = append(coins, coin)
	}

	for i := n - 1; i >= 0; i-- {
		now = coins[i]
		if k/now >= 1 {
			answer += k / now
			k %= now
		}
		if k == 0 {
			break
		}
	}
	fmt.Println(answer)
}
