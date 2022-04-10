package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	c := (a * b) % 2
	if c == 0 {
		fmt.Printf("Even")
	} else {
		fmt.Printf("Odd")
	}
}
