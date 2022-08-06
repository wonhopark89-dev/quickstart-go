package main

import (
	"fmt"
	"tutorial_go/something"
)

func main() {
	fmt.Println("main")
	something.SayHello()

	const name string = "John"
	// name = "Jane" 불가
	fmt.Println(name)
	
	var name2 string = "john"
	name2 = "jane"
	fmt.Println(name2)

	name3 := "john"
	name3 = "jane"
	fmt.Println(name3)
}