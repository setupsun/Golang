package main

import (
	"fmt"
)

func getUserInput() (string, string, string) {
	var name, city, hobby string
	
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	
	fmt.Print("Enter city: ")
	fmt.Scanln(&city)
	
	fmt.Print("Enter hobby: ")
	fmt.Scanln(&hobby)

	return name, city, hobby
}

func stringCreate(name string, city string, hobby string) string {
	return fmt.Sprintf("Hi! My name %s. I'm from %s. My hobby is %s.", name, city, hobby)
}

func main() {
	name, city, hobby := getUserInput()

	message1 := stringCreate("Jack", "New York", "kvadrobing")
	fmt.Println("Old version without user input: ", message1)

	message2 := stringCreate(name, city, hobby)
	fmt.Println("New version: ", message2)
}