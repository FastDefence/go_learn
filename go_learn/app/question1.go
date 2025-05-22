package main

import (
	"fmt"
	"math"
)

func sqrt(n int) float64 {
	return math.Sqrt((float64(n)))
}

func pow(m int, n int) int {
	return int(math.Pow(float64(m), float64(n)))
}

func main() {
	var l int
	_, err1 := fmt.Scan(&l)
	if err1 != nil {
		fmt.Println("読み取りエラー:", err1)
		return
	}
	fmt.Println(sqrt(l))
	var m, n int
	_, err2 := fmt.Scan(&m, &n)
	if err2 != nil {
		fmt.Println("読み取りエラー:", err2)
		return
	}
	fmt.Println(pow(m, n))
}
