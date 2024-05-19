package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/stock/entity"
	avltree "github.com/vltvdnl/Go-Stock-Scrapper.git/pkg/avl_tree"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/pkg/postgres"
)

type StockRepo struct {
	*postgres.Postgres             // DB
	*avltree.AVLTree[entity.Stock] // RAM
}

func New(pg *postgres.Postgres) *StockRepo {
	sqlstatement := `CREATE TABLE IF NOT EXISTS stocks(
		id serial PRIMARY KEY,
		name text,
		symb text unique,
		price real,
		curchange real,
		change_in_per real
	);`
	_, err := pg.DB.Exec(sqlstatement)
	if err != nil {
		log.Fatalf("internal - stock - repo - New: %v", err)
	}
	return &StockRepo{pg, &avltree.AVLTree[entity.Stock]{}}
}

func (r *StockRepo) GetAll(ctx context.Context) ([]entity.Stock, error) {
	if r.AVLTree.FindMax() == nil {
		log.Println("From DB")
		sqlstatement := `SELECT name, symb, price, curchange, change_in_per  FROM stocks;`
		rows, err := r.DB.QueryContext(ctx, sqlstatement)
		if err != nil {
			return nil, fmt.Errorf("internal - stock - usecase - repo - stock_postgres - GetALL: %v", err)
		}
		defer rows.Close()
		entities := make([]entity.Stock, 0, 100)
		for rows.Next() {
			e := entity.Stock{}
			err := rows.Scan(&e.Name, &e.Symb, &e.Price, &e.CurChange, &e.ChangeInPer)
			if err != nil {
				return nil, fmt.Errorf("internal - stock - usecase - repo - stock_postgres - GetAll: %v", err)
			}
			r.AVLTree.Insert(int(e.Price), e)
			entities = append(entities, e)
		}
		return entities, nil
	}
	log.Println("From RAM")
	entities := r.AVLTree.ToSlice()
	return entities, nil
}

func (r *StockRepo) Store(ctx context.Context, s entity.Stock) error {
	sqlstatement := `INSERT INTO stocks(name, symb, price, curchange, change_in_per) VALUES ($1, $2, $3, $4, $5) 
	ON CONFLICT (symb) DO UPDATE SET price = $3, curchange = $4, change_in_per = $5;`
	go func() {
		_, err := r.DB.ExecContext(ctx, sqlstatement, s.Name, s.Symb, s.Price, s.CurChange, s.ChangeInPer)
		if err != nil {
			log.Printf("internal - stock - usecase - repo - stock_postgres - Store: %v", err)
			return
		}
	}()
	r.AVLTree.Insert(int(s.Price), s)
	return nil
}
func (r *StockRepo) FindMin(ctx context.Context) (*entity.Stock, error) {
	sqlstatement := `SELECT 
	name, symb, price, curchange, change_in_per 
	FROM stocks where price = (SELECT MIN(price) FROM stocks);`
	stock := r.AVLTree.FindMin()
	if stock == nil {
		log.Println("From DB")
		var s entity.Stock
		row := r.DB.QueryRowContext(ctx, sqlstatement)
		err := row.Scan(&s.Name, &s.Symb, &s.Price, &s.CurChange, &s.ChangeInPer)
		if err != nil {
			return nil, fmt.Errorf("error in FindMin")
		}
		return &s, nil

	}
	log.Println("From RAM")
	return stock, nil

}
func (r *StockRepo) FindMax(ctx context.Context) (*entity.Stock, error) {
	sqlstatement := `SELECT name, symb, price, curchange, change_in_per
	 FROM stocks where price = (SELECT MAX(price) FROM stocks);`
	stock := r.AVLTree.FindMax()
	if stock == nil {
		var s entity.Stock
		row := r.DB.QueryRowContext(ctx, sqlstatement)
		err := row.Scan(&s.Name, &s.Symb, &s.Price, &s.CurChange, &s.ChangeInPer)
		if err != nil {
			return nil, fmt.Errorf("error in FindMAX")
		}
		return &s, nil

	}
	return stock, nil

}
