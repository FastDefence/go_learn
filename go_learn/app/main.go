package main

import (
	"fmt"
)

func main() {
	var m, n int

	fmt.Println("平方根を計算します。")
	fmt.Scan(&n)
	fmt.Println(sqrt(n))

	fmt.Println("累乗を計算します。")
	fmt.Scan(&m, &n)
	fmt.Println(pow(m, n))

	fmt.Println("九九を計算します。")
	fmt.Scan(&m, &n)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", times(i, j))
		}
		fmt.Println()
	}
}
