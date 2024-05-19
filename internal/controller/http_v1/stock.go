package httpv1

import (
	"fmt"
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
		h.GET("/min", r.min)
		h.GET("/max", r.max)
	}
}

type stocksResponse struct {
	Stocks []entity.Stock `json:"stocks"`
}
type stockResponse struct {
	Stock entity.Stock `json:"stock"`
}

func (r *stockRoutes) stocks(c *gin.Context) {
	stocks, err := r.s.AllStocks(c.Request.Context())
	if err != nil {
		c.Error(fmt.Errorf("No min whyyyy %v", err))
		return
	}
	c.JSON(http.StatusOK, stocksResponse{stocks})
}
func (r *stockRoutes) min(c *gin.Context) {
	stock, err := r.s.MinStock(c.Request.Context())
	if err != nil {
		c.Error(fmt.Errorf("No min whyyyy %v", err))
		return
	}
	c.JSON(http.StatusOK, stockResponse{*stock})
}
func (r *stockRoutes) max(c *gin.Context) {
	stock, err := r.s.MaxStock(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, stockResponse{*stock})
}
