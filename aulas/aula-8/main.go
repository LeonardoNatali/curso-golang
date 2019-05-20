package main

import "fmt"

func main() {
	fatNum := 25
	soma(12, 14, 17, 19)
	fmt.Printf("Fatorial de %d = %d", fatNum, fatorial(fatNum))
}

func soma(params ...int) {
	for key, value := range params {
		fmt.Printf("[RANGE] params[%d]: %d\n", key, value)
	}
}

func fatorial(num int) int {
	var res int
	if num > 0 {
		if num > 1 {
			res = num * fatorial(num-1)
			fmt.Printf("Num: %d - Fat: %d\n", num, res)
		} else {
			res = num
		}
	} else {
		res = 0
	}
	return res
}
