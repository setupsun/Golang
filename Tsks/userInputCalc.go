package main

import (
	"fmt"
)

func getUserInputData() (float64, float64, string) {
	var value1, value2 float64
	var operation string

	fmt.Println("Enter first value: ")
	fmt.Scanln(&value1)

	fmt.Println("Enter second value: ")
	fmt.Scanln(&value2)

	fmt.Println("Enter operation: ")
	fmt.Scanln(&operation)

	return value1, value2, operation
}

func calculateData(value1 float64, value2 float64, operation string) string {
	var answer float64

	switch operation {
	case "+":
		answer = value1 + value2
	case "-":
		answer = value1 - value2
	case "*":
		answer = value1 * value2
	case "/":
		if value2 != 0 {
			answer = value1 / value2
		} else {
			return "Devided by zero is prohibited"
		}
	case ":":
		if value2 != 0 {
			answer = value1 / value2
		} else {
			return "Devided by zero is prohibited"
		}
	default:
		return "Choose a valid operation! List of operations: +, -, *, /"
	}
	return fmt.Sprintf("= %g", answer)
}

func main() {
	value1, value2, operation := getUserInputData()

	message := calculateData(value1, value2, operation)
	fmt.Println("Answer : ", message)
}
