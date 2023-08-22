package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// var baseUrl = "https://kr.iherb.com/search?kw=%EB%B9%84%ED%83%80%EB%AF%BC&"

type extractedItem struct {
	id    string
	title string
}

// Scrape iherb by a term
func Scrape(term string) {
	var baseUrl = "https://kr.iherb.com/search?kw=" + term + "&"
	var items []extractedItem
	c := make(chan []extractedItem)
	totalPage := getPages(baseUrl)

	for i := 1; i <= totalPage; i++ {
		// items = append(items, getPage(i)...)
		go getPage(i, baseUrl, c)
	}

	for i := 0; i < totalPage; i++ {
		extractedItems := <-c
		items = append(items, extractedItems...)
	}

	// fmt.Println(items)
	for _, item := range items {
		fmt.Println(item.id, item.title)
	}

	// write file
	writeItems(items)
	fmt.Println("Done, extracted", len(items))
}

func writeItems(items []extractedItem) {
	file, err := os.Create("items.csv")
	checkErr(err)

	// 함수가 끝나는 시점에 파일에 데이터를 입력하는 함수
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, item := range items {
		itemSlice := []string{item.id, item.title}
		itemErr := w.Write(itemSlice)
		checkErr(itemErr)
	}

}

func getPage(page int, url string, mainChannel chan<- []extractedItem) {
	var items []extractedItem
	c := make(chan extractedItem)

	pageURL := url + "p=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchItems := doc.Find(".product-cell-container")
	searchItems.Each(func(i int, item *goquery.Selection) {
		// items = append(items, extractItem(item, c))
		go extractItem(item, c)
	})

	for i := 0; i < searchItems.Length(); i++ {
		item := <-c
		items = append(items, item)
	}

	// return items
	mainChannel <- items
}

func extractItem(item *goquery.Selection, c chan<- extractedItem) {
	id, _ := item.Find(".product").Attr("itemid")

	// title := item.Find(".product-inner>div").Each(func(i int, item *goquery.Selection) {
	// 	content, exist := item.Find(".product-title").Attr("content")
	// 	if exist {
	// 		// fmt.Println(id, content)
	// 		c <- extractedItem{
	// 			id:    id,
	// 			title: content,
	// 	}
	// )}

	title, _ := item.Find(".product-inner>div").Find(".product-title").Attr("content")

	c <- extractedItem{
		id:    id,
		title: title,
	}
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
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
