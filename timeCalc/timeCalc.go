//Задача: Калькулятор времени
//Напиши программу, которая:
//
//Запрашивает у пользователя два момента времени в формате HH:MM (например, "14:30" и "18:45").
//Рассчитывает разницу между этими моментами в минутах и часах.
//Выводит результат в формате:
//makefile
//Copy code
//Разница: 4 часа 15 минут.

package main

import (
	"fmt"
	"time"
)

func getUserInput() string {
	var inputTime string
	fmt.Print("Enter time (HH:MM or 'Now'): ")
	fmt.Scan(&inputTime)
	return inputTime
}

func parseTime(timeStr string) (time.Time, error) {
	if timeStr == "Now" {
		return time.Now(), nil
	}

	// Парсим только время (HH:MM), привязывая его к текущей дате
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, err
	}

	// Привязываем время к сегодняшней дате
	now := time.Now()
	parsedTime = time.Date(now.Year(), now.Month(), now.Day(),
		parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())
	return parsedTime, nil
}

func calculateTimeDifference(time1, time2 time.Time) (int, int) {
	diff := time2.Sub(time1)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) % 60
	return hours, minutes
}

func main() {
	fmt.Println("Enter the first time:")
	timeStr1 := getUserInput()

	fmt.Println("Enter the second time:")
	timeStr2 := getUserInput()

	parsedTime1, err1 := parseTime(timeStr1)
	parsedTime2, err2 := parseTime(timeStr2)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input time. Please use the HH:MM format or 'Now'.")
		return
	}

	// Вычисляем разницу
	hours, minutes := calculateTimeDifference(parsedTime1, parsedTime2)

	// Корректный вывод
	fmt.Printf("Time difference: %d hours and %d minutes.\n", hours, minutes)
}
