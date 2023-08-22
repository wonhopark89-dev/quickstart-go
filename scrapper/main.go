package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strings"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(c.FormValue("term"))
	fmt.Println(term)
	fmt.Println(len(term))
	return nil
}

func main() {
	// scrapper.Scrape("vita")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

}
