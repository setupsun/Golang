//Проверка на чётность
// Напиши программу, которая:
// Запрашивает у пользователя число.
// Определяет, является ли оно чётным или нечётным.
// Выводит результат в формате:
// "Число X является чётным"
// "Число X является нечётным"


package main

import (
	"fmt"
	"math"
	"strings"
)

func getUserNum() float64 {
	var number float64
	fmt.Println("Enter number (Example : 1; 1.0, 1.3) : ")
	fmt.Scanln(&number)
	return number
}

func countDecimalPlaces(number float64) int {
	getStringFromNumber := fmt.Sprintf("%f", number)
	if dotIndex := strings.Index(getStringFromNumber, "."); dotIndex != -1 {
		getStringFromNumber = strings.TrimRight(getStringFromNumber, "0")
		return len(getStringFromNumber[dotIndex+1:])
	}
	return 0
}

func calcParity(number float64, numberPower int) string {
	// Получаем целую часть числа
	intPart := int(number)
	// Проверяем чётность целой части
	var result string
	if intPart%2 == 0 {
		result = fmt.Sprintf("Целая часть %d чётная", intPart)
	} else {
		result = fmt.Sprintf("Целая часть %d нечётная", intPart)
	}

	// Обрабатываем дробную часть
	numberMultiplicator := int(math.Pow(float64(10), float64(numberPower)))
	resultNumber := int(number * float64(numberMultiplicator))
	if resultNumber%2 == 0 {
		result += fmt.Sprintf(", дробная часть чётная")
	} else {
		result += fmt.Sprintf(", дробная часть нечётная")
	}

	return result
}

func main() {
	number := getUserNum()
	numberPower := countDecimalPlaces(number)
	message := calcParity(number, numberPower)
	fmt.Println(message)
}
