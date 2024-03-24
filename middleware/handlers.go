package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/storage"
)

func GetAllStock(w http.ResponseWriter, r *http.Request) { //
	w.Header().Set("Content-Type", "application/json")
	stocks, err := storage.DB_GetAllStock()
	if err != nil {
		fmt.Println("Unable to load stocks form db") // need logger
		return
	}
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
	stock, err := storage.DB_GetSpecStock(name)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Some troubles with name (maybe you do some error)")
		return
	}
	if stock.Name == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Fuck you")
		return
	}
	json.NewEncoder(w).Encode(stock)
}

func GetAllCoins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	coins, err := storage.DB_GetAllCoins()
	if err != nil {
		fmt.Println("Unable to load coins from db")
		json.NewEncoder(w).Encode("Unable to load coins from db")
		return
	}
	if len(coins) == 0 {
		log.Println("Unable to get cryptocurreny")
		return
	}
	json.NewEncoder(w).Encode(coins)
}

func GetCoin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := mux.Vars(r)
	coin, err := storage.DB_GetSpecCoin(name["name"])
	if err != nil {
		fmt.Println("Unable to load coin")
	}
	if coin.Name == "" {
		fmt.Println("Unable to load coin from db")
		json.NewEncoder(w).Encode("Unable to load coin from db")
		return
	}
	json.NewEncoder(w).Encode(coin)
}
