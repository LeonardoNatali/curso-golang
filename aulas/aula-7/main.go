package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	user := map[string]string{
		"name":      "Leonardo",
		"last_name": "Natali",
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("[FOR] arr[%d]: %d\n", i, arr[i])
	}

	for key, value := range arr {
		fmt.Printf("[RANGE] arr[%d]: %d\n", key, value)
	}

	for key, value := range user {
		fmt.Printf("[RANGE] map[%s]: %s\n", key, value)
	}
}
