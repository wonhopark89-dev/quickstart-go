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

func superAdd(numbers ...int) int {
	total := 0
	
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	
	// range 활용
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

func canIDrink(age int) bool {
  // basic if 
  // if age < 18 {
  //   return false
  // }   
  // return true

  // advanced if
  // sampleAge := age + 3; // 함수 안에서 사용할 수 있음
  if koreanAge := age + 2; koreanAge < 18 {
    // koreanAge 는 if 문 안에서만 사용할 수 있음
    return false
  }  
  return true
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

	total := superAdd(1, 2, 3, 4, 5)
	fmt.Println("sum:", total)

	// repeatMe("aa","bb","cc")

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

  fmt.Println(canIDrink(16))
}