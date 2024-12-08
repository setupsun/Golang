package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Получение строки от пользователя
func getUserInput() string {
	fmt.Print("Enter text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Подсчёт количества символов
func lengthOfString(input string) int {
	return len(input)
}

// Подсчёт количества слов
func countingWords(input string) int {
	words := strings.Fields(input)
	return len(words)
}

// Подсчёт уникальных слов
func countingUniqueWords(input string) int {
	processed := preprocessString(input)
	words := strings.Fields(processed)
	wordSet := make(map[string]struct{})
	for _, word := range words {
		wordSet[word] = struct{}{}
	}
	return len(wordSet)
}

// Поиск самого длинного слова
func longestWord(input string) string {
	processed := preprocessString(input)
	words := strings.Fields(processed)
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// Обработка строки: удаление знаков препинания и приведение к нижнему регистру
func preprocessString(input string) string {
	processed := strings.ToLower(input)
	processed = strings.ReplaceAll(processed, ".", "")
	processed = strings.ReplaceAll(processed, ",", "")
	return processed
}

func main() {
	input := getUserInput()

	fmt.Printf("Number of characters: %d\n", lengthOfString(input))
	fmt.Printf("Number of words: %d\n", countingWords(input))
	fmt.Printf("Number of unique words: %d\n", countingUniqueWords(input))
	fmt.Printf("Longest word: %s\n", longestWord(input))
}
