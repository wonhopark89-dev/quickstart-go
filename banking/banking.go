package main

import (
	"banking/accounts"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	account := accounts.NewAccount("pepe")
	fmt.Println(account)
}
