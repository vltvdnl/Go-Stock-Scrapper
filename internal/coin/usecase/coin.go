package usecase

import (
	"context"
	"fmt"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/entity"
)

type CoinUseCase struct {
	repo   CoinRepo
	webAPI CoinWebAPI
}

func New(r CoinRepo, w CoinWebAPI) *CoinUseCase {
	return &CoinUseCase{repo: r, webAPI: w}
}

func (u *CoinUseCase) AllCoins(ctx context.Context) ([]entity.Coin, error) {
	coins, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("internal - coin - usecase - u.AllCoins - repo.GetAll: %v", err)
	}
	return coins, nil
}

func (u *CoinUseCase) GetCoins(ctx context.Context) error {
	coins, err := u.webAPI.GetCoins()
	if err != nil {
		return fmt.Errorf("internal - coin - usecase - u.AllCoins - repo.GetCoins: %v", err)
	}
	for _, coin := range coins {
		err = u.repo.Store(ctx, coin)
		if err != nil {
			return fmt.Errorf("internal - coin - usecase - u.AllCoins - repo.GetAll: %v", err)
		}
	}
	return nil
}
