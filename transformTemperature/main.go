//Задача: Конвертер температур
//Напиши программу, которая:
//
//Принимает от пользователя значение температуры.
//Запрашивает единицу измерения (C — градусы Цельсия, F — градусы Фаренгейта, K — Кельвины).
//Предлагает выбрать, в какую единицу перевести температуру.
//Выполняет пересчёт и выводит результат.

package main

import (
	"fmt"
)

func getUserInput() (string, string, float64) {
	var unit, targetUnit string
	var degrees float64

	validUnits := map[string]bool{"C": true, "F": true, "K": true}

	fmt.Println("Enter unit (C - degrees Celsius, F - degrees Fahrenheit, K - Kelvin): ")
	fmt.Scanln(&unit)

	fmt.Println("Enter target unit (C - degrees Celsius, F - degrees Fahrenheit, K - Kelvin): ")
	fmt.Scanln(&targetUnit)

	if !validUnits[unit] || !validUnits[targetUnit] {
		fmt.Println("Error: Invalid unit. Please use C, F, or K.")
		return "", "", 0
	}

	fmt.Println("Enter degrees (e.g., 1.00): ")
	_, err := fmt.Scanln(&degrees)
	if err != nil {
		fmt.Println("Error reading degrees:", err)
		return "", "", 0
	}

	if unit == "K" && degrees < 0 {
		fmt.Println("Error: Kelvin temperature cannot be less than 0.")
		return "", "", 0
	}

	return unit, targetUnit, degrees
}

func transformTemp(degrees float64, unit string, targetUnit string) float64 {
	var result float64

	switch unit {
	case "C":
		switch targetUnit {
		case "F":
			result = degrees*9/5 + 32
		case "K":
			result = degrees + 273.15
		}
	case "F":
		switch targetUnit {
		case "C":
			result = (degrees - 32) * 5 / 9
		case "K":
			result = (degrees-32)*5/9 + 273.15
		}
	case "K":
		switch targetUnit {
		case "C":
			result = degrees - 273.15
		case "F":
			result = (degrees-273.15)*9/5 + 32
		}
	}
	return result
}

func main() {
	unit, targetUnit, degrees := getUserInput()
	if unit == "" || targetUnit == "" {
		return
	}

	message := transformTemp(degrees, unit, targetUnit)
	fmt.Printf("Result: %.2f %s\n", message, targetUnit)
}
