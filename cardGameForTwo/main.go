package main

import (
	"fmt"
	"sort"
)

// sliceを高順にソート
func main() {
	var n, alice, bob int
	fmt.Scanf("%d", &n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &cards[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	for i := 0; i < n; i += 2 {
		alice += cards[i]
		if i+1 < n {
			bob += cards[i+1]
		}
	}
	fmt.Println(alice - bob)
}
