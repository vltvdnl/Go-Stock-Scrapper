package webapi

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/entity"
)

type CoinWebAPI struct{}

func New() *CoinWebAPI {
	return &CoinWebAPI{}
}

func (w *CoinWebAPI) GetCoins() ([]entity.Coin, error) {
	coins := []entity.Coin{}

	col := colly.NewCollector()

	col.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	col.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	col.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		var err error
		coin := entity.Coin{}

		coin.Name = e.ChildText(".cmc-table__cell--sort-by__name")
		coin.Symb = e.ChildText(".cmc-table__cell--sort-by__symbol")
		if len(e.ChildText(".cmc-table__cell--sort-by__price")) > 0 {
			coin.Price, err = strconv.ParseFloat(strings.ReplaceAll(e.ChildText(".cmc-table__cell--sort-by__price")[1:], ",", ""), 64)
			if err != nil {
				log.Fatalf("error while parsing: %v", err) // пофиксить потом
			}
		} else {
			coin.Price = 0.0
		}

		coin.HourChangePer, err = strconv.ParseFloat(strings.ReplaceAll(e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"), "%", ""), 64)
		if err != nil {
			coin.HourChangePer = 0.0
			log.Printf("error while parsing: %v", err) // пофиксить потом
		}
		coin.DayChangePer, err = strconv.ParseFloat(strings.ReplaceAll(e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"), "%", ""), 64)
		if err != nil {
			coin.DayChangePer = 0.0
			log.Printf("error while parsing: %v", err) // пофиксить потом
		}
		coin.WeekChangePer, err = strconv.ParseFloat(strings.ReplaceAll(e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"), "%", ""), 64)
		if err != nil {
			coin.WeekChangePer = 0.0
			log.Printf("error while parsing: %v", err) // пофиксить потом
		}
		coins = append(coins, coin)
	})
	col.Visit("https://coinmarketcap.com/all/views/all")
	// log.Println(coins)
	return coins, nil
}
