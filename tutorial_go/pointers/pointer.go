package main

import "fmt"

func main() {
	a := 2
	b := a
	c := &a
	a = 10
	*c = 20 // a 는 20으로 업데이트 됨 
	fmt.Println(a,b)
	fmt.Println(&a, &b) // memory address
	fmt.Println(c, *c) // pointer value, *c
	
} 