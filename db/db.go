package db

import (
	"database/sql"
	"fmt"
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

func (db *DB) Signup(user, password, firstName, lastName, bio string, birthday time.Time) error {
	query := `call signUp_Procudure(?, ?, ?, ?, ?, ?)`
	_, err := db.SqlDB.Exec(query, user, password, firstName, lastName, birthday, bio)
	return err
}

func (db *DB) Login(user, password string) error {
	query := `call logIn_Procedure(?, ?)`
	res, _ := db.SqlDB.Exec(query, user, password)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("invalid username or password")
	}
	return err
}
