package main

import (
	"github.com/labstack/echo/v4"
	"os"
	"scrapper/scrapper"
	"strings"
)

const fileName string = "items.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName) // 서버(폴더)에 있는 파일 삭제하기 위해 defer 사용 ( 다운로드 후 삭제 )
	term := strings.ToLower(c.FormValue("term"))
	// fmt.Println(term)
	// fmt.Println(len(term))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	// scrapper.Scrape("vita")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

}
