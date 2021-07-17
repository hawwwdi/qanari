package model

import (
	"database/sql"
	"encoding/json"
)

type Ava struct {
	ID          int           `json:"id" binding:"-"`
	Publisher   string        `json:"publisher" binding:"-"`
	Content     string        `json:"content" binding:"-"`
	ReplyTo     sql.NullInt32 `json:"replyTo" binding:"-"`
	PublishTime string        `json:"publishTime" binding:"-"`
	Likes       int           `json:"likes" binding:"-"`
}

func (a Ava) MarshalJSON() ([]byte, error) {
	res := make(map[string]interface{})
	res["id"] = a.ID
	res["publisher"] = a.Publisher
	res["content"] = a.Content
	res["publishTime"] = a.PublishTime
	res["likes"] = a.Likes
	if a.ReplyTo.Valid {
		res["replyTo"] = a.ReplyTo.Int32
	}
	return json.Marshal(res)
}

func (a *Ava) UnmarshalJSON(b []byte) error {
	res := make(map[string]interface{})
	_ = json.Unmarshal(b, &res)
	if _, valid := res["id"]; valid {
		a.ID = int(res["id"].(float64))
	}
	if _, valid := res["publisher"]; valid {
		a.Publisher = res["publisher"].(string)
	}
	if _, valid := res["content"]; valid {
		a.Content = res["content"].(string)
	}
	if _, valid := res["publishTime"]; valid {
		a.PublishTime = res["publishTime"].(string)
	}
	if _, valid := res["likes"]; valid {
		a.Likes = int(res["likes"].(float64))
	}
	if _, valid := res["replyTo"]; valid {
		a.ReplyTo.Int32 = int32(res["replyTo"].(float64))
		a.ReplyTo.Valid = true
	}
	return nil
}
