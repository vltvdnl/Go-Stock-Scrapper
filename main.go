package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	name   string
	price  string
	change string
}

func main() {
	// ticker := []string{
	// 	"aflt",
	// 	"vtbr",
	// 	"gazp",
	// 	"lkoh",
	// 	"yndx",
	// }

	file, err := os.Create("stocks.csv")

	if err != nil {
		log.Fatalln("Failed to create a file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"name",
		"price",
		"change",
	}

	writer.Write(headers)

	defer writer.Flush()

	col := colly.NewCollector()
	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	col.OnError(func(_ *colly.Response, err error) {
		log.Println("Error ", err)
	})

	col.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		stock := Stock{}
		stock.name = e.ChildText(".data-col1")
		stock.price = e.ChildText(".data-col2")
		stock.change = e.ChildText(".data-col4")
		fmt.Println("Company name: ", stock.name)
		fmt.Println("Price: ", stock.price)
		fmt.Println("Change: ", stock.change)

		line := []string{
			stock.name,
			stock.price,
			stock.change,
		}
		writer.Write(line)

	})
	// col.Wait()

	col.Visit("https://finance.yahoo.com/lookup")

	// fmt.Println(stocks)

}
