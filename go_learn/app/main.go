package main

import (
	"fmt"
)

func main() {
	var m, n int
	//q1
	fmt.Println("平方根を計算します。")
	fmt.Scan(&n)
	fmt.Println(sqrt(n))

	fmt.Println("累乗を計算します。")
	fmt.Scan(&m, &n)
	fmt.Println(pow(m, n))

	//q2
	fmt.Println("九九を計算します。")
	fmt.Scan(&m, &n)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", times(i, j))
		}
		fmt.Println()
	}

	//q3
	fmt.Println("テストの点数")
	fmt.Scan(&n)
	iftest(n)

	fmt.Println("テストの点数")
	fmt.Scan(&n)
	swicthtest(n)

	//q4
	numbers := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var i, j int
	for i = 0; i < len(numbers); i++ {
		for j = i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	fmt.Println(numbers)

	//q5
	imap := map[string]string{"apple": "120", "banana": "100", "orange": "150"}
	for name, value := range imap {
		fmt.Printf("%s: %s\n", name, value)
	}
}
