package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	cryptscrap "github.com/vltvdnl/Go-Stock-Scrapper.git/CryptScrap"
	stockscrap "github.com/vltvdnl/Go-Stock-Scrapper.git/StockScrap"
)

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stocks := stockscrap.AllStock()
	if len(stocks) == 0 {
		log.Println("Unable to get stoks")
	}
	json.NewEncoder(w).Encode(stocks)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	l := mux.Vars(r)
	name := l["name"]
	fmt.Println(name)
	stock := stockscrap.RequestStock(name)
	if stock.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("you")
		return
	}
	json.NewEncoder(w).Encode(stock)
}

func GetAllCoins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	coins := cryptscrap.AllCrypts()
	if len(coins) == 0 {
		log.Println("Unable to get cryptocurreny")
	}
	json.NewEncoder(w).Encode(coins)
}

func GetCoin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := mux.Vars(r)
	coin := cryptscrap.RequestCoin(name["name"])
	if coin.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Fuck you")
		return
	}
	json.NewEncoder(w).Encode(coin)
}
