package storage

// TODO: error lib
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	cryptscrap "github.com/vltvdnl/Go-Stock-Scrapper.git/CryptScrap"
	stockscrap "github.com/vltvdnl/Go-Stock-Scrapper.git/StockScrap"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func CreateConnection() *sql.DB {
	psqlconnection := fmt.Sprintf("host = %s port = %d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconnection)

	if err != nil {
		log.Fatal("Connection to DB failed: ", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("No answer from db: ", err)
	}

	return db
}

// TODO:
// 1) Download all data to db
// 2) Get all data from db
// 3)
func DB_PutStocks(Stocks []stockscrap.Stock) {
	db := CreateConnection()
	defer db.Close()
	sqlState := `INSERT INTO stocks(symb, name, price, curchange, perchange) VALUES($1, $2, $3, $4, $5)
	ON CONFLICT (symb) DO UPDATE SET price=$3, curchange=$4, perchange=$5;` // могут быть проблемы с запросом, надо смотреть ON CONFLICT (symb) DO UPDATE SET (price, curchange, perchange) VALUES($3, $4, $5)

	for _, val := range Stocks {
		_, err := db.Exec(sqlState,
			val.Symb,
			val.Name,
			val.Price,
			val.CurChange,
			val.PerChange)

		if err != nil {
			log.Fatalf("Unable to put data (stocks) to DB: %v", err)
		}
	}
	log.Println("Stocks are inserted") // maybe show time when it's done
}
func DB_PutCoins(Coins []cryptscrap.Crypto) {
	db := CreateConnection()
	defer db.Close()
	sqlState := `INSERT INTO coins(rank, name, symb, price, hourchange, daychange, weekchange) VALUES($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (name) DO UPDATE SET price=$4, hourchange=$5, daychange=$6, weekchange=$7;`
	for _, val := range Coins {
		_, err := db.Exec(sqlState,
			val.Rank,
			val.Name,
			val.Symb,
			val.Price,
			val.HourChangePer,
			val.DayChangePer,
			val.WeekChangePer)
		if err != nil {
			log.Fatalf("Unable to put data to DB: %v", err) //
		}
	}

	log.Println("Coins are inserted")
}
func DB_GetAllStock() ([]stockscrap.Stock, error) { // maybe stocks and coins in one function
	db := CreateConnection()
	defer db.Close()
	var Stocks []stockscrap.Stock
	sqlState := `SELECT * FROM stocks`
	rows, err := db.Query(sqlState)
	if err != nil {
		log.Fatal("Unable to execute query") // not fatal later ??
	}
	defer rows.Close()
	for rows.Next() {
		var stock stockscrap.Stock
		err = rows.Scan(&stock.Symb,
			&stock.Name,
			&stock.Price,
			&stock.CurChange,
			&stock.PerChange)
		if err != nil {
			log.Fatal("Unable to scan a row")
		}
		Stocks = append(Stocks, stock)
	}
	return Stocks, nil
}
func DB_GetAllCoins() ([]cryptscrap.Crypto, error) {
	var Coins []cryptscrap.Crypto
	db := CreateConnection()
	defer db.Close()
	sqlState := `SELECT * FROM coins`
	rows, err := db.Query(sqlState)

	if err != nil {
		log.Fatal("Unable to execute query")
	}
	defer rows.Close()
	for rows.Next() {
		var coin cryptscrap.Crypto
		err = rows.Scan(
			&coin.Rank,
			&coin.Name,
			&coin.Symb,
			&coin.Price,
			&coin.HourChangePer,
			&coin.DayChangePer,
			&coin.WeekChangePer)
		if err != nil {
			log.Fatalf("Unable to scan a row: %v", err)
		}
		Coins = append(Coins, coin)
	}
	return Coins, err
}
func DB_GetSpecStock(name_or_rank string) (stockscrap.Stock, error) { // не работает блять надо смотреть
	db := CreateConnection()
	defer db.Close()
	var stock stockscrap.Stock

	sqlState := `SELECT * FROM stocks WHERE symb=$1 OR name=$1`

	row := db.QueryRow(sqlState, name_or_rank)
	err := row.Scan(&stock.Symb,
		&stock.Name,
		&stock.Price,
		&stock.CurChange,
		&stock.PerChange)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan a row: %v", err)
		// return _, err
	}
	return stock, err
}
func DB_GetSpecCoin(name_or_rank string) (cryptscrap.Crypto, error) {
	db := CreateConnection()
	defer db.Close()
	var coin cryptscrap.Crypto
	sqlState := `SELECT * FROM coins WHERE name=$1 OR symb=$1`

	row := db.QueryRow(sqlState, name_or_rank)

	err := row.Scan(&coin.Rank,
		&coin.Name,
		&coin.Symb,
		&coin.Price,
		&coin.HourChangePer,
		&coin.DayChangePer,
		&coin.WeekChangePer)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No were returned")
		return coin, nil
	case nil:
		return coin, nil
	default:
		log.Printf("Unable to scan a row: %v \n", err)
	}
	return coin, err
}
