package httpv1

import (
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
	}
}

type coinsResponse struct {
	Coins []entity.Coin `json:"coins"`
}

func (r *coinRoutes) coins(c *gin.Context) {
	coins, err := r.c.AllCoins(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coinsResponse{coins})
}
