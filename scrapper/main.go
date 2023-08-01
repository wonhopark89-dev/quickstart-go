package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

var baseUrl string = "https://kr.iherb.com/new-products?"

func main() {
	totalPage := getPages()

	for i := 1; i <= totalPage; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseUrl + "p=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		maxLength := s.Find("a").Length()
		// fmt.Println(maxLength)
		pages = maxLength
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status code:", res.StatusCode)
	}
}
