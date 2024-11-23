package main

import (
	"fmt"
)

func stringCreate(name string, city string, hobby string) string {
	return fmt.Sprintf("Hi! My name %s . I'm from %s . My hobby is %s", name, city, hobby)
}

func main() {
	message := stringCreate("Jack", "New York", "kvadrobing")
	fmt.Println(message)
}