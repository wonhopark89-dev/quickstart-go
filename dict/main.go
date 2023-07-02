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

	bye := "bye"
	dictionary.Add(bye, "See you")
	err3 := dictionary.Update(bye, "Goodbye")
	if err3 != nil {
		fmt.Println(err3)
	}
	word, _ := dictionary.Search(bye)
	fmt.Println(word)

	dictionary.Delete(bye)
	_, err4 := dictionary.Search(bye)
	if err4 != nil {
		fmt.Println(err4)
	}
}
