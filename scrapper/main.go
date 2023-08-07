package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

var baseUrl string = "https://kr.iherb.com/new-products?"

type extractedItem struct {
	id    string
	title string
}

func main() {
	var items []extractedItem
	totalPage := getPages()

	for i := 1; i <= totalPage; i++ {
		items = append(items, getPage(i)...)
	}

	// fmt.Println(items)
	for _, item := range items {
		fmt.Println(item.id, item.title)
	}
}

func getPage(page int) []extractedItem {
	var items []extractedItem
	pageURL := baseUrl + "p=" + strconv.Itoa(page)
	// fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".product-cell-container").Each(func(i int, item *goquery.Selection) {
		id, _ := item.Find(".product").Attr("itemid")

		item.Find(".product-inner>div").Each(func(i int, item *goquery.Selection) {
			content, exist := item.Find(".product-title").Attr("content")
			if exist {
				// fmt.Println(id, content)
				items = append(items, extractedItem{
					id:    id,
					title: content,
				})
			}
		})
		// fmt.Println(id)
	})
	return items
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
