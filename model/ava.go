package model

import "time"

type Ava struct {
	ID          int64     `json:"id" binding:"-"`
	Publisher   int64     `json:"publisher" binding:"-"`
	Content     string    `json:"content" binding:"required"`
	ReplyTo     int64     `json:"replyTo" binding:"-"`
	PublishTime time.Time `json:"publishTime" binding:"-"`
}
