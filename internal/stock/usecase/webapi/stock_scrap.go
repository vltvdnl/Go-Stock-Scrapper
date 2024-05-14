package webapi

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/entity"
)

type StockWebAPI struct{}

func New() *StockWebAPI {
	return &StockWebAPI{}
}

func (w *StockWebAPI) GetStocks() ([]entity.Stock, error) {
	stocks := make([]entity.Stock, 0, 102)
	url := launcher.New().
		Headless(true).
		Devtools(false).
		MustLaunch()
	browser := rod.New().
		ControlURL(url).
		MustConnect().
		NoDefaultDevice()
	defer browser.MustClose()
	log.Println("Visiting: https://www.cnbc.com/nasdaq-100/")
	page := browser.MustPage("https://www.cnbc.com/nasdaq-100/")
	// time.Sleep(2 * time.Second)
	TextFromXpath := func(xpath string) string {
		name, err := page.ElementX(xpath)
		if err != nil {
			log.Fatalf("some trouble: %v", err)
		}
		text, err := name.Text()
		if err != nil {
			log.Fatalf("some trouble with text: %v", err)
		}
		return text
	}
	for i := 1; i <= 101; i++ {
		var stock entity.Stock
		var err error

		symbXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[1]/div/div/a", i)
		nameXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[2]", i)
		priceXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[3]", i)
		curchangeXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[4]", i)
		perchangeXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[5]", i)

		stock.Symb = TextFromXpath(symbXpath)
		stock.Name = TextFromXpath(nameXpath)
		stock.Price, err = strconv.ParseFloat(strings.ReplaceAll(TextFromXpath(priceXpath), ",", ""), 64)
		if err != nil {
			return nil, fmt.Errorf("internal - stock - usecase - webapi - GetStocks: %v", err)
		}

		stock.CurChange, err = strconv.ParseFloat(TextFromXpath(curchangeXpath), 64)
		if err != nil {
			stock.CurChange = 0.0
		}
		stock.ChangeInPer, err = strconv.ParseFloat(TextFromXpath(perchangeXpath), 64)
		if err != nil {
			stock.ChangeInPer = 0.0
		}

		stocks = append(stocks, stock)
	}
	log.Println("Query is complete")
	return stocks, nil

}
