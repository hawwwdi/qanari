package model

import (
	"database/sql"
	"encoding/json"
)

type Message struct {
	ID       int            `json:"id" binding:"-"`
	Sender   int            `json:"sender" binding:"-"`
	Receiver string         `json:"receiver" binding:"required"`
	Content  sql.NullString `json:"content" binding:"-"`
	PostID   sql.NullInt32  `json:"postID" binding:"-"`
	SendTime string         `json:"sendTime" binding:"-"`
}

func (m Message) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})
	res["id"] = m.ID
	res["sender"] = m.Sender
	res["receiver"] = m.Receiver
	res["sendTime"] = m.SendTime
	if m.Content.Valid {
		res["content"] = m.Content.String
	}
	if m.PostID.Valid {
		res["postID"] = m.PostID.Int32
	}
	return json.Marshal(res)
}

func (m *Message) UnmarshalJSON(b []byte) error {
	res := make(map[string]interface{})
	_ = json.Unmarshal(b, &res)
	if _, valid := res["id"]; valid {
		m.ID = int(res["id"].(float64))
	}
	if _, valid := res["sender"]; valid {
		m.Sender = int(res["sender"].(float64))
	}
	if _, valid := res["receiver"]; valid {
		m.Receiver = res["receiver"].(string)
	}
	if _, valid := res["sendTime"]; valid {
		m.SendTime = res["sendTime"].(string)
	}
	if _, valid := res["content"]; valid {
		m.Content.String = res["content"].(string)
		m.Content.Valid = true
	}
	if _, valid := res["postID"]; valid {
		m.PostID.Int32 = int32(res["postID"].(float64))
		m.PostID.Valid = true
	}
	return nil
}
