package stockscrap

// TODO: too slow solution with go-rod, maybe need to return to gocolly but i dont't now how to do it))
import (
	"fmt"
	"log"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
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
		var stock Stock
		symbXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[1]/div/div/a", i)
		nameXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[2]", i)
		priceXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[3]", i)
		curchangeXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[4]", i)
		perchangeXpath := fmt.Sprintf("/html/body/div[3]/div/div[1]/div[3]/div/div/div/div[3]/div[1]/section/section/div/div/div/div[1]/div/table/tbody/tr[%d]/td[5]", i)

		stock.Symb = TextFromXpath(symbXpath)

		stock.Name = TextFromXpath(nameXpath)

		stock.Price = TextFromXpath(priceXpath)
		stock.Price = strings.ReplaceAll(stock.Price, ",", "")

		stock.CurChange = TextFromXpath(curchangeXpath)
		if stock.CurChange == "UNCH" {
			stock.CurChange = "0"
		}
		stock.PerChange = TextFromXpath(perchangeXpath)
		if stock.PerChange == "UNCH" {
			stock.PerChange = "0"
		}

		stocks = append(stocks, stock)
	}
	return stocks

	// col := colly.NewCollector()

	// col.OnError(func(_ *colly.Response, err error) {
	// 	log.Println("Error: ", err)
	// })

	// col.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting: ", r.URL)
	// })

	// col.OnHTML("tbody tr", func(e *colly.HTMLElement) {
	// 	stock := Stock{}
	// 	stock.Symb = e.ChildText(".Va(m) Ta(end) Pstart(20px) Fw(600) Fz(s)")
	// 	stock.Name = e.ChildText(".Name")
	// 	stock.Price = e.ChildText(".Price (Intraday)")
	// 	stock.CurChange = e.ChildText(".Change")
	// 	stock.PerChange = e.ChildText(".% Change")
	// 	stocks = append(stocks, stock)
	// 	fmt.Println(stock) // отладка
	// })
	// col.Visit("https://finance.yahoo.com/most-active/") // done (not tested)
	// // пиздец какой-то блять всё переделывать ...
	// return stocks
}
