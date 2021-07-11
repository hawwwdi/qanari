package model

type Ava struct {
	ID          int64  `json:"id" binding:"-"`
	Publisher   int64  `json:"publisher" binding:"-"`
	Content     string `json:"content" binding:"-"`
	ReplyTo     int64  `json:"replyTo" binding:"-"`
	PublishTime string `json:"publishTime" binding:"-"`
	Likes       int    `json:"likes" binding:"-"`
}
