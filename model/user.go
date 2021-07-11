package model

type SignupCredential struct {
	ID        int    `json:"id" binding:"-"`
	Username  string `json:"Username" binding:"required"`
	Password  string `json:"Password" binding:"required"`
	FirstName string `json:"FirstName" binding:"required"`
	LastName  string `json:"LastName" binding:"required"`
	Bio       string `json:"Bio" binding:"-"`
	Birthday  string `json:"Birthday" binding:"required"`
}

type LoginCredential struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type User struct {
	ID       int64  `json:"Id" binding:"-"`
	Username string `json:"Username" binding:"required"`
}
