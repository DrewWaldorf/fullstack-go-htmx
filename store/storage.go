package store

import (
	"database/sql"
	//"log"
	"fmt"
	//"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Imonaboat1."
	dbname   = "postgres"
)

func NewMySQLStorage() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	return db, nil
}
