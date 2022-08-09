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

// 리턴 타입 
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
// 리턴 타입 및 지정 ( naked return )
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("lenAndUpper2 done ...")
	length = len(name)
	uppercase = strings.ToUpper(name)
	// return length, uppercase 아래처럼 생략 가능 naked
	return
}


func repeatMe(words ...string) {
	fmt.Println(words)
}
 
func main() {
	fmt.Println("main")

	fmt.Println(multiply(2, 3))

	totalLength, upperName := lenAndUpper("hello123")
	fmt.Println(totalLength, upperName)

	totalLength2, _ := lenAndUpper("bye123")
	fmt.Println(totalLength2)

	totalLength3, upperName3 := lenAndUpper2("bye456")
	fmt.Println(totalLength3, upperName3)


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