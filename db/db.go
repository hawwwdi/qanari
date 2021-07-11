package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hawwwdi/qanari/model"
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

func (db *DB) Signup(s *model.SignupCredential) error {
	query := `call signUp_Procudure(?, ?, ?, ?, ?, ?)`
	_, err := db.SqlDB.Exec(query, s.Username, s.Password, s.FirstName, s.LastName, s.Bio, s.Birthday)
	return err
}

func (db *DB) Login(l *model.LoginCredential) error {
	query := `call logIn_Procedure(?, ?)`
	res, _ := db.SqlDB.Exec(query, l.Username, l.Password)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("invalid username or password")
	}
	return err
}

func (db *DB) PostAva(a *model.Ava) error {
	query := `call postAVA_procedure(?)`
	_, err := db.SqlDB.Exec(query, a.Content)
	return err
}

func (db *DB) PostComment(a *model.Ava) error {
	query := `call postComment_procedure(?, ?)`
	_, err := db.SqlDB.Exec(query, a.Content, a.ReplyTo)
	return err
}
