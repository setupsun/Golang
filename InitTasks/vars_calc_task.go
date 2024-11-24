package main 

import "fmt"

func calcAge(year int, currentYear int) string{
	calcData := currentYear - year 
	// return "Ваш возраст: " calcData
	return fmt.Sprintf("Ваш возраст: %d", calcData)
}

func main() {
	message := calcAge(1999, 2024)
	fmt.Println(message)
}