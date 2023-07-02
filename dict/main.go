package main

import (
	"dict/mydict"
	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err2 := dictionary.Add("first", "Greeting")
	if err2 != nil {
		fmt.Println(err2)
	}
}
