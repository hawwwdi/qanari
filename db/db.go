package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SqlDB *sql.DB
}

func NewDB(user, password, dbName string) *DB {
	db, err := sql.Open("mysql", user+":"+password+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 30)
	return &DB{db}
}
