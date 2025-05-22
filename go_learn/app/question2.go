package main

import (
	"fmt"
)

func times(m int, n int) {

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%3d * %3d = %3d\n", i, j, i*j)
		}
	}
}
