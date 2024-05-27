package database

import (
	"database/sql"

	"github.com/gustafer/go-games/cmd/api/configs"
	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	connStr := configs.GetConnString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate() error {
	connStr := configs.GetConnString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	sql := `CREATE TABLE IF NOT EXISTS games (
		id SERIAL PRIMARY KEY,
		title varchar(45) NOT NULL,
		description varchar(450) NOT NULL
		);`
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
