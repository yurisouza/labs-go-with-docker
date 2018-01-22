package repositories

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	errDB error
	once  sync.Once
)

const (
	host     = "172.18.0.2"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func GetInstanceDB() (*sql.DB, error) {

	once.Do(func() {
		db, _ = configInstanceDB()
	})

	if p := db.Ping(); p != nil {
		err := p
		return nil, err
	}

	return db, nil
}

func configInstanceDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, errDB = sql.Open("postgres", psqlInfo)
	return db, errDB
}
