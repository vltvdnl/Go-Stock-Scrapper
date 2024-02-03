package stockscrap

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	Symb      string
	Name      string
	Price     string
	CurChange string
	PerChange string
}

func AllStock() []Stock {
	stocks := []Stock{}

	col := colly.NewCollector()

	col.OnError(func(_ *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	col.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	col.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.Symb = e.ChildText(".data-col0")
		stock.Name = e.ChildText(".data-col1")
		stock.Price = e.ChildText(".data-col2")
		stock.CurChange = e.ChildText(".data-col3")
		stock.PerChange = e.ChildText("data-col4")
		stocks = append(stocks, stock)
	})
	col.Visit("https://finance.yahoo.com/lookup")
	return stocks
}
func RequestStock(name_symb string) Stock {
	log.Println(name_symb)
	stocks := AllStock()
	for _, stock := range stocks {
		if stock.Name == name_symb || stock.Symb == name_symb {
			return stock
		}
	}
	log.Println("No match in stocks")
	return Stock{}
}
