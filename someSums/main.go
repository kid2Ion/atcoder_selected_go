package main

import "fmt"

// 格桁の足し算関数
func sumOfDegit(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var n, a, b, result int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	for i := 1; i < n+1; i++ {
		if sum := sumOfDegit(i); a <= sum && sum <= b {
			result += i
		}
	}
	fmt.Println(result)
}
