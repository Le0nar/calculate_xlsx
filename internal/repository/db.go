package repository

import (
	"database/sql"

	_ "github.com/lib/pq" // подключаем драйвер PostgreSQL
)

func NewDB() (*sql.DB, error) {
	connStr := "user=postgres password=qwerty dbname=postgres host=localhost port=5436 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
