package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/vltvdnl/Go-Stock-Scrapper.git/internal/coin/entity"
	avltree "github.com/vltvdnl/Go-Stock-Scrapper.git/pkg/avl_tree"
	"github.com/vltvdnl/Go-Stock-Scrapper.git/pkg/postgres"
)

type CoinRepo struct {
	*postgres.Postgres
	*avltree.AVLTree[entity.Coin]
}

func New(pg *postgres.Postgres) *CoinRepo {
	sqlstatement := `CREATE TABLE IF NOT EXISTS coins(
		id serial PRIMARY KEY,
		name text,
		symb text UNIQUE,
		price real,
		hourchange_per real,
		daychange_per real,
		weekchange_per real
	);`
	_, err := pg.DB.Exec(sqlstatement)
	if err != nil {
		log.Fatalf("internal - coin - usecase - repo - New: %v", err)
	}
	return &CoinRepo{pg, &avltree.AVLTree[entity.Coin]{}}
}

func (r *CoinRepo) GetAll(ctx context.Context) ([]entity.Coin, error) {
	entities := make([]entity.Coin, 0, 150)
	if r.AVLTree.FindMax() == nil {
		sqlstatement := `SELECT name, symb, price, hourchange_per, daychange_per, weekchange_per FROM coins`
		rows, err := r.DB.QueryContext(ctx, sqlstatement)
		if err != nil {
			// log.Printf("internal - usecase - repo - coin_postgres: %v", err)
			return nil, fmt.Errorf("internal - usecase - repo - coin_postgres - GetAll: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			e := entity.Coin{}

			err := rows.Scan(&e.Name, &e.Symb, &e.Price, &e.HourChangePer, &e.DayChangePer, &e.WeekChangePer)

			if err != nil {
				return nil, fmt.Errorf("internal - usecase - repo - coin_postges: %v", err)
			}
			entities = append(entities, e)
		}
		return entities, nil
	}
	entities = r.AVLTree.ToSlice()
	return entities, nil
}

func (r *CoinRepo) Store(ctx context.Context, c entity.Coin) error {
	go func() {
		sqlstatement := `INSERT INTO coins(name, symb, price, hourchange_per, daychange_per, weekchange_per) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT(symb) DO UPDATE SET price = $3, hourchange_per = $4, daychange_per = $5, weekchange_per =$6;`
		_, err := r.DB.ExecContext(ctx, sqlstatement, c.Name, c.Symb, c.Price, c.HourChangePer, c.DayChangePer, c.WeekChangePer)
		if err != nil {
			log.Printf("internal - usercase - repo - coin_postgres - Store: %v", err)
			return
		}
	}()
	r.AVLTree.Insert(int(c.Price), c)
	return nil
}
func (r *CoinRepo) FindMax(ctx context.Context) (*entity.Coin, error) {
	sqlstatement := `SELECT 
	name, symb, price, hourchange_per, daychange_per, weekchange_per
	FROM coins where price = (SELECT MAX(price) FROM coins);`
	stock := r.AVLTree.FindMax()
	if stock == nil {
		var s entity.Coin
		row := r.DB.QueryRowContext(ctx, sqlstatement)
		err := row.Scan(&s.Name, &s.Symb, &s.Price, &s.HourChangePer, &s.DayChangePer, &s.WeekChangePer)
		if err != nil {
			return nil, fmt.Errorf("error in FindMAX")
		}
		return &s, nil

	}
	return stock, nil
}
func (r *CoinRepo) FindMin(ctx context.Context) (*entity.Coin, error) {
	sqlstatement := `SELECT 
	name, symb, price, hourchange_per, daychange_per, weekchange_per
	FROM coins where price = (SELECT MIN(price) FROM coins);`
	stock := r.AVLTree.FindMin()
	if stock == nil {
		var s entity.Coin
		row := r.DB.QueryRowContext(ctx, sqlstatement)
		err := row.Scan(&s.Name, &s.Symb, &s.Price, &s.HourChangePer, &s.DayChangePer, &s.WeekChangePer)
		if err != nil {
			return nil, fmt.Errorf("error in FindMin")
		}
		return &s, nil

	}
	return stock, nil
}
