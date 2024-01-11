package main

import "fmt"

func Calculator(values ...int) {
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Printf("The result of addition : %d\n", result)

	result = 0
	for i, v := range values {
		if i == 0 {
			result += v
		} else {
			result *= v
		}
	}
	fmt.Printf("The result of multiplication : %d\n", result)
}
