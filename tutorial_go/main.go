package main

import (
	"fmt"
	"strings"
	"tutorial_go/something"
)

// func multiply(a, int, b int) int {}
// 축약
func multiply(x, y int) int {
	return x * y
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
 
func main() {
	fmt.Println("main")

	fmt.Println(multiply(2, 3))

	totalLength, upperName := lenAndUpper("hello123");
	fmt.Println(totalLength, upperName)

	totalLength2, _ := lenAndUpper("bye123");
	fmt.Println(totalLength2)


	repeatMe("aa","bb","cc")

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