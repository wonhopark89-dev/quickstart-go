package main

import "fmt"

type Person struct {
	name string
	age  int
	food []string
}

func main() {
	food := []string{"apple", "banana", "orange"}
	p1 := Person{"John", 20, food}
	p2 := Person{name: "Paul", age: 30, food: food}
	fmt.Println(p1)
	fmt.Println(p2.food)
}
