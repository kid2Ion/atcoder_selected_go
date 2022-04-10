package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int
	var total int
	fmt.Scanf("%d", &a)
	astr := strconv.Itoa(a)
	nums := strings.Split(astr, "")

	for _, num := range nums {
		if num == "1" {
			total++
		}
	}
	fmt.Println(total)
}
