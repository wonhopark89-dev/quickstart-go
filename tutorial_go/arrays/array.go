package main

import "fmt"

func main() {
	// array
	names := [3]string{"John", "Paul", "George"}
	// append(names, "Ringo")
	// names[3] = "Ringo" // out of index
	fmt.Println(names)

	// slice
	animals := []string{"dog", "cat", "bird"}
	animals = append(animals, "fish") // 새로운 배열 리턴
	// animals[10] = "horse" // out of index
	fmt.Println(animals)
}