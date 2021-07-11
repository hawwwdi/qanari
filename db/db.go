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

func (db *DB) Like(a *model.Ava) error {
	query := `call likeAva_procedure(?)`
	res, err := db.SqlDB.Exec(query, a.ID)
	rows, err := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("Ava not found")
	}
	return err
}

func (db *DB) SendMessage(m *model.Message) error {
	query := `call sendMessage(?, ?, ?)`
	res, err := db.SqlDB.Exec(query, m.Receiver, m.Content, m.PostID)
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("message not send")
	}
	return err
}

func (db *DB) GetTimeLine() ([]*model.Ava, error) {
	query := `call getTimeLine_procedure()`
	rows, err := db.SqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timeLine := make([]*model.Ava, 1)
	for rows.Next() {
		var curr model.Ava
		if err := rows.Scan(&curr.ID, &curr.Publisher, &curr.Content, &curr.ReplyTo, &curr.PublishTime); err != nil {
			return nil, err
		}
		timeLine = append(timeLine, &curr)
	}
	return timeLine, nil
}
