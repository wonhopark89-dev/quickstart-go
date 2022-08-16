package main

import "fmt"

func main() {
	// map
	// map[key]value, key :string, value :int
	ages := map[string]int{
		"John":  20,
		"Paul":  30,
		"George": 40,
	}
	ages["Ringo"] = 50
	fmt.Println(ages)
	
	for key ,value := range ages {
		fmt.Println(key, value)
	}

	// delete
	delete(ages, "George")
	fmt.Println(ages)
}