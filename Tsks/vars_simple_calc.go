package main

import (
	"fmt"
)

func simpCalc(var1 float64, var2 float64, mathSign string) string {

	if mathSign == "+" {
		answer := var1 + var2
		return fmt.Sprintf("Answer: %g", answer)
	} else if mathSign == "-" {
		answer := var1 - var2
		return fmt.Sprintf("Answer: %g", answer)
	} else if mathSign == "*" {
		answer := var1 * var2
		return fmt.Sprintf("Answer: %g", answer)
	} else if mathSign == "/" {
		answer := var1 / var2
		return fmt.Sprintf("Answer: %g", answer)
	} else {
		answer := "Choose correct operation! List of operations : +, -, *, /"
		return fmt.Sprintf(answer)
	}
}

func updatedSimpleCalc(var1 float64, var2 float64, mathSign string) string {
	var answer float64

	switch mathSign{
	case "+":
		answer = var1 + var2
	case "-":
		answer = var1 - var2
	case "*":
		answer = var1 * var2
	case "/":
		if var2 != 0{
			answer = var1 / var2
		}else {
			return "Devided by zero is prohibited"
		}
	default:
		return "Choose a valid operation! List of operations: +, -, *, /"
	}
	return fmt.Sprintf("= %g",answer)
}

func main() {
	message1 := simpCalc(8.6, 2.5, "*")
	fmt.Println(message1)
	message2 := updatedSimpleCalc(8.6, 0, "%")
	fmt.Println("Your solution: ", message2)
}