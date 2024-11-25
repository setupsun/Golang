// Задача:
// Напиши программу, которая принимает от пользователя строку и проверяет,
// является ли она палиндромом (строка, которая читается одинаково слева направо
// 	и справа налево, игнорируя пробелы, знаки препинания и регистр).

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func getUserInput() string {
	var userInput string
	fmt.Print("Enter string: ")
	fmt.Scanln(&userInput)
	return userInput
}

func cleanString(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	userInput := getUserInput()
	cleanedInput := cleanString(userInput)
	reversedInput := reverseString(cleanedInput)

	if cleanedInput == reversedInput {
		fmt.Println("The input is a palindrome.")
	} else {
		fmt.Println("The input is not a palindrome.")
	}
}
