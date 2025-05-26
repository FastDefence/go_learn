package main

func iftest(n int) {
	if n >= 0 && n < 60 {
		println("不合格")
	} else if n < 80 {
		println("合格")
	} else if n <= 100 {
		println("大金星")
	} else {
		println("？")
	}
}

func swicthtest(n int) {
	switch {
	case n >= 0 && n < 60:
		println("不合格")
	case n < 80:
		println("合格")
	case n <= 100:
		println("大金星")
	default:
		println("？")
	}
}
