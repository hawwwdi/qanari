package db

import (
	"fmt"

	"github.com/hawwwdi/qanari/model"
)

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

func (db *DB) Follow(u *model.User) error {
	query := `call follow(?)`
	res, err := db.SqlDB.Exec(query, u.Username)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user %s already followed", u.Username)
	}
	return err
}

func (db *DB) UnFollow(u *model.User) error {
	query := `call unfollow(?)`
	res, err := db.SqlDB.Exec(query, u.Username)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user %s not found in following list", u.Username)
	}
	return err
}

func (db *DB) Block(u *model.User) error {
	query := `call block_procedure(?)`
	res, err := db.SqlDB.Exec(query, u.Username)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user %s already blocked or not found", u.Username)
	}
	return err
}

func (db *DB) UnBlock(u *model.User) error {
	query := `call unblock_procedure(?)`
	res, err := db.SqlDB.Exec(query, u.Username)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("user %s not found in blocked users list", u.Username)
	}
	return err
}
