package juggling_araray

import "fmt"

//최대공약수 gcd를 이용해 집합을 나누어 여러 요소를 한꺼번에 이동 시키는 것

//최대 공약수 gcd 활용
func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func leftRotate(arr []int, d int, n int) {
	i := 0
	temp := arr[i]
	for i = 0; i < gcd(d, n); i++ {
		j := i
		for true {
			k := j + d
			if k >= n {
				k = k - n
			}
			if k == i {
				break
			}
			arr[j] = arr[k]
			j = k
		}
		arr[j] = temp
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(arr)
	fmt.Println(arr)
	leftRotate(arr, 2, n)
	fmt.Println(arr)
}
