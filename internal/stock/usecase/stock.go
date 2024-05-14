package usecase

import (
	"context"
	"fmt"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/entity"
)

type StockUseCase struct {
	repo   StockRepo
	webAPI StockWebAPI
}

func New(r StockRepo, w StockWebAPI) *StockUseCase {
	return &StockUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (u *StockUseCase) AllStocks(ctx context.Context) ([]entity.Stock, error) {
	stocks, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("internal - stock - usecase - StockUseCase - u.repo.GetAll: %w", err)
	}
	return stocks, nil
}
func (u *StockUseCase) GetStocks(ctx context.Context) error {
	stocks, err := u.webAPI.GetStocks()
	if err != nil {
		return fmt.Errorf("internal - stock - usecase - StockUseCase - u.repo.GetStocks: %w", err)
	}
	for _, stock := range stocks {
		err = u.repo.Store(ctx, stock)
		if err != nil {
			return fmt.Errorf("internal - stock - usecase - StockUseCase - u.repo.GetStocks: %w", err)
		}
	}
	return nil
}
