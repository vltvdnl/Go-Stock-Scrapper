package router

import (
	"github.com/gorilla/mux"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/stocks", middleware.GetAllStock).Methods("GET")
	r.HandleFunc("/stocks/{name}", middleware.GetStock).Methods("GET")
	r.HandleFunc("/crypt", middleware.GetAllCoins).Methods("GET")
	r.HandleFunc("/crypt/{name}", middleware.GetCoin).Methods("GET")
	return r
}
