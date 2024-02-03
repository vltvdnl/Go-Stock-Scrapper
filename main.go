package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starring server on port 8080: ...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
