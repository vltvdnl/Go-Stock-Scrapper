package usecase

import (
	"context"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/entity"
)

type (
	Stocks interface {
		AllStocks(ctx context.Context) ([]entity.Stock, error)
	}
	StockRepo interface {
		GetAll(ctx context.Context) ([]entity.Stock, error)
		Store(ctx context.Context, s entity.Stock) error
	}

	StockWebAPI interface {
		GetStocks() ([]entity.Stock, error)
	}
)
