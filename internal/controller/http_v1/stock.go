package httpv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/entity"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/usecase"
)

type stockRoutes struct {
	s usecase.Stocks
}

func newStocksRoutes(handler *gin.RouterGroup, s usecase.Stocks) {
	r := &stockRoutes{s}

	h := handler.Group("/stocks")
	{
		h.GET("/all", r.stocks)
	}
}

type stocksResponse struct {
	Stocks []entity.Stock `json:"stocks"`
}

func (r *stockRoutes) stocks(c *gin.Context) {
	stocks, err := r.s.AllStocks(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, stocksResponse{stocks})
}
