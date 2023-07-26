package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	//  make: 빈배열 생성
	// var results = make(map[string]string)
	// urls := []string{
	// 	"https://www.airbnb.com/",
	// 	"https://www.google.com/",
	// 	"https://www.amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://www.google.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.facebook.com/",
	// 	"https://www.instagram.com/",
	// }

	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitUrl(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// }

	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }

	c := make(chan bool)
	greetings := [2]string{"hello", "bye"}
	for _, greeting := range greetings {
		go isGreeting(greeting, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	// go helloCount("hello")
	// go helloCount("bye")
	// time.Sleep(time.Second * 5)
}

func hitUrl(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

func helloCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is count", i)
		time.Sleep(time.Second)
	}
}

func isGreeting(person string, ch chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person, ch)
	ch <- true
}
