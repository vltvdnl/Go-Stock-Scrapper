package httpv1

import (
	"github.com/gin-gonic/gin"
	coin "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/usecase"
	stock "github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/usecase"
)

func NewRouter(s stock.Stocks, c coin.Coins) *gin.Engine {
	handler := gin.New()

	h := handler.Group("/v1")
	{
		newCoinRoutes(h, c)
		newStocksRoutes(h, s)
	}
	return handler
}
