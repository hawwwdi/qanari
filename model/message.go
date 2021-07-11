package model

type Message struct {
	ID       int    `json:"ID" binding:"-"`
	Sender   int    `json:"Sender" binding:"-"`
	Receiver int    `json:"Receiver" binding:"required"`
	Content  string `json:"Content" binding:"-"`
	PostID   int    `json:"PostID" binding:"-"`
	SendTime string `json:"SendTime" binding:"-"`
}
