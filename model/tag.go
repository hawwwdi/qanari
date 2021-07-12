package model

type Tag struct {
	ID  int    `json:"id" binding:"-"`
	Tag string `json:"tag" binding:"required"`
}
