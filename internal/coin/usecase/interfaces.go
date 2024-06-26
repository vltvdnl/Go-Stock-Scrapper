package usecase

import (
	"context"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/entity"
)

type (
	Coins interface {
		AllCoins(ctx context.Context) ([]entity.Coin, error)
		MinCoin(ctx context.Context) (*entity.Coin, error)
		MaxCoin(ctx context.Context) (*entity.Coin, error)
	}

	CoinRepo interface {
		GetAll(ctx context.Context) ([]entity.Coin, error)
		Store(ctx context.Context, c entity.Coin) error
		FindMax(ctx context.Context) (*entity.Coin, error)
		FindMin(ctx context.Context) (*entity.Coin, error)
	}
	CoinWebAPI interface {
		GetCoins() ([]entity.Coin, error)
	}
)
