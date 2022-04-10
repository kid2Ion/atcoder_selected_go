package main

// 重複を配列から省く方法→連想配列
import "fmt"

func main() {
	var n, mochi int
	fmt.Scanf("%d", &n)
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &mochi)
		m[mochi]++
	}

	fmt.Println(len(m))
}
