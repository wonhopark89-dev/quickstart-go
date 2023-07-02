package main

import (
	"banking/accounts"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello")
	account := accounts.NewAccount("pepe")
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
