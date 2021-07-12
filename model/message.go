package model

type Message struct {
	ID       int    `json:"id" binding:"-"`
	Sender   int    `json:"sender" binding:"-"`
	Receiver int    `json:"receiver" binding:"required"`
	Content  string `json:"content" binding:"-"`
	PostID   int    `json:"postID" binding:"-"`
	SendTime string `json:"sendTime" binding:"-"`
}
