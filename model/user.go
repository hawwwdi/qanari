package model

type SignupCredential struct {
	ID        int    `json:"id" binding:"-"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Bio       string `json:"bio" binding:"-"`
	Birthday  string `json:"birthday" binding:"required"`
}

type LoginCredential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int64  `json:"id" binding:"-"`
	Username string `json:"username" binding:"required"`
}
