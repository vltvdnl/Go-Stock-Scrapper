package cryptscrap

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Crypto struct {
	Rank          string
	Name          string
	Symb          string
	Price         string
	HourChangePer string
	DayChangePer  string
	WeekChangePer string
}

func AllCrypts() []Crypto {
	coins := []Crypto{}

	col := colly.NewCollector()

	col.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	col.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	col.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		coin := Crypto{}

		coin.Rank = e.ChildText(".cmc-table__cell--sort-by__rank")
		if coin.Rank == "" {
			return
		}
		coin.Name = e.ChildText(".cmc-table__cell--sort-by__name")
		coin.Symb = e.ChildText(".cmc-table__cell--sort-by__symbol")
		coin.Price = e.ChildText(".cmc-table__cell--sort-by__price")[1:]
		coin.Price = strings.ReplaceAll(coin.Price, ",", "")
		coin.HourChangePer = e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h")
		coin.DayChangePer = e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h")
		coin.WeekChangePer = e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d")
		coins = append(coins, coin)
	})
	col.Visit("https://coinmarketcap.com/all/views/all")
	// log.Println(coins)
	return coins
}

// func RequestCoin(name_symb string) Crypto {
// 	coins := AllCrypts()
// 	for _, coin := range coins {
// 		if coin.Name == name_symb || coin.Symb == name_symb {
// 			return coin
// 		}
// 	}
// 	log.Println("No match in cryptocoins")
// 	return Crypto{}
// }
