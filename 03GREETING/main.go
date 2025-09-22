package main

import (
	"fmt"
)

func main() {
	name := "Aditya"
	age := 19
	msg := fmt.Sprintf("My name is %s and i'm %d years old", name, age)
	fmt.Printf(msg)
}
