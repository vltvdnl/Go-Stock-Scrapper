package main

import (
	"fmt"
	"log"
	"net/http"

	cryptscrap "github.com/vltvdnl/Go-Stock-Scrapper.git/CryptScrap"
	stockscrap "github.com/vltvdnl/Go-Stock-Scrapper.git/StockScrap"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/router"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/storage"
)

// потоки легли нахуй друг с другом... бля походу надо табы переелывать))))
func main() {
	go storage.DB_PutStocks(stockscrap.AllStock())
	go storage.DB_PutCoins(cryptscrap.AllCrypts())
	r := router.Router()
	fmt.Println("Starring server on port 8080: ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
