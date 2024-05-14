package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func New(url string) (*Postgres, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("pkg - postgres - New: Can't connect to postgres server: %v", err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatalf("pkg - postgres - New: No answer from postgres server: %v", err)
	}
	return &Postgres{DB: db}, nil
}
