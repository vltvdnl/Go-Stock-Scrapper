package httpv1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/entity"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/usecase"
)

type coinRoutes struct {
	c usecase.Coins
}

func newCoinRoutes(handler *gin.RouterGroup, c usecase.Coins) {
	r := &coinRoutes{c}
	h := handler.Group("/coins")
	{
		h.GET("/all", r.coins)
		h.GET("/min", r.min)
		h.GET("/max", r.max)
	}
}

type coinsResponse struct {
	Coins []entity.Coin `json:"coins"`
}
type coinResponse struct {
	Stock entity.Coin `json:"coin"`
}

func (r *coinRoutes) coins(c *gin.Context) {
	coins, err := r.c.AllCoins(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coinsResponse{coins})
}
func (r *coinRoutes) min(c *gin.Context) {
	stock, err := r.c.MinCoin(c.Request.Context())
	if err != nil {
		c.Error(fmt.Errorf("No min whyyyy %v", err))
		return
	}
	c.JSON(http.StatusOK, coinResponse{*stock})
}
func (r *coinRoutes) max(c *gin.Context) {
	stock, err := r.c.MaxCoin(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coinResponse{*stock})
}
