package main

import (
	"fmt"
)

func main() {
	var m, n int

	fmt.Println("平方根を計算します。")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("読み取りエラー:", err)
		return
	}
	fmt.Println(sqrt(n))

	fmt.Println("累乗を計算します。")
	fmt.Scan(&m, &n)
	if err != nil {
		fmt.Println("読み取りエラー:", err)
		return
	}
	fmt.Println(pow(m, n))

	fmt.Println("九九を計算します。")
	fmt.Scan(&m, &n)
	kuku(m, n)
}
