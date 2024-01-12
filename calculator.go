package go_calculator

func Calculator(values ...int) (addition, multiply int) {
	result := 0
	for _, v := range values {
		result += v
	}
	addition = result

	result = 0
	for i, v := range values {
		if i == 0 {
			result += v
		} else {
			result *= v
		}
	}
	multiply = result

	return
}
