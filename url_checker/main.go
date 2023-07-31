package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("request failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	// make: 빈배열 생성
	// var results map[string]string = make(map[string]string)
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}

	//c := make(chan string)
	//greetings := [5]string{"hello", "bye", "ccc", "ddd", "eee"}
	//for _, greeting := range greetings {
	//	go isGreeting(greeting, c)
	//}

}

// send only
func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

func helloCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is count", i)
		time.Sleep(time.Second)
	}
}

func isGreeting(person string, ch chan string) {
	time.Sleep(time.Second * 5)
	// fmt.Println(person, ch)
	ch <- person + " is greeting"
}
