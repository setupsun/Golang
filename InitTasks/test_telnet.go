package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализация рандома

	// Бесконечный цикл
	for {
		// Генерация случайных значений
		ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		port := rand.Intn(65535)        // Случайный порт
		protocol := []string{"TCP", "UDP"}[rand.Intn(2)] // Случайный протокол
		status := []string{"CONNECTED", "DISCONNECTED", "TIMEOUT"}[rand.Intn(3)] // Случайный статус

		// Вывод строки
		fmt.Printf("Telnet log: IP=%s PORT=%d PROTOCOL=%s STATUS=%s\n", ip, port, protocol, status)

		// Задержка перед следующей итерацией
		time.Sleep(1 * time.Nanosecond) // Задержка 1 секунда
	}
}
