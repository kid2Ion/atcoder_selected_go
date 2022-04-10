package main

import "fmt"

func main() {
	var n int
	count := -1
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	flg := true
	for flg {
		for i := 0; i < n; i++ {
			if nums[i]%2 == 0 {
				nums[i] /= 2
			} else {
				flg = false
			}
		}
		count++
	}
	fmt.Println(count)
}
