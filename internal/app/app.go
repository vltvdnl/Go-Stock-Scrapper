package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/config"
	coin_usecase "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/usecase"
	coin_repo "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/usecase/repo"
	coin_webapi "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/usecase/webapi"
	httpv1 "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/controller/http_v1"
	stock_usecase "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/usecase"
	stock_repo "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/usecase/repo"
	stock_webapi "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/usecase/webapi"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/pkg/httpserver"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/pkg/postgres"
)

func Run(cfg *config.Config) {

	pg, err := postgres.New(cfg.PG.String())
	if err != nil {
		log.Fatalf("app - Run - postgres.New: %v", err)
	}
	defer pg.DB.Close()

	stockUseCase := stock_usecase.New(
		stock_repo.New(pg),
		stock_webapi.New(),
	)

	go func() {
		err = stockUseCase.GetStocks(context.TODO())
		if err != nil {
			log.Printf("error in stocks api: %v", err)
		}
	}()

	coinUseCase := coin_usecase.New(
		coin_repo.New(pg),
		coin_webapi.New(),
	)

	go func() {
		err = coinUseCase.GetCoins(context.TODO())
		if err != nil {
			log.Printf("error in coins api: %v", err)
		}
	}()

	httpServer := httpserver.New(httpv1.NewRouter(stockUseCase, coinUseCase), cfg.HTTP.Port)
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println(s.String())
	case err = <-httpServer.Notify():
		log.Printf("error: %v", err)
	}

}
