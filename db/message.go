package db

import (
	"fmt"

	"github.com/hawwwdi/qanari/model"
)

func (db *DB) SendMessage(m model.Message) error {
	query := `call sendMessage(?, ?, ?)`
	res, err := db.SqlDB.Exec(query, m.Receiver, m.Content, m.PostID)
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("message not send")
	}
	return err
}

func (db *DB) GetMessagesFromAUser(u model.User) ([]model.Message, error) {
	query := `call getMessagesFromAUser(?)`
	rows, err := db.SqlDB.Query(query, u.Username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	messages := make([]model.Message, 0)
	for rows.Next() {
		var curr model.Message
		if err := rows.Scan(&curr.ID, &curr.Sender, &curr.Receiver, &curr.Content, &curr.PostID, &curr.SendTime); err != nil {
			return nil, err
		}
		messages = append(messages, curr)
	}
	return messages, nil
}

func (db *DB) GetMessageSenders() ([]model.User, error) {
	query := `call getSenders()`
	rows, err := db.SqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	senders := make([]model.User, 0)
	for rows.Next() {
		var curr model.User
		if err := rows.Scan(&curr.ID, &curr.Username); err != nil {
			return nil, err
		}
		senders = append(senders, curr)
	}
	return senders, nil
}
