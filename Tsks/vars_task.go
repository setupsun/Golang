package main

import "fmt"

// vars in-interface
func greet(name string) string {
	return "Привет, " + name + "!"
}

func main() {
	message := greet("Саша")
	fmt.Println(message)
}
