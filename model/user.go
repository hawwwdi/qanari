package model

import "time"

type SignupCredential struct {
	Username  string    `json:"Username" binding:"required"`
	Password  string    `json:"Password" binding:"required"`
	FirstName string    `json:"FirstName" binding:"required"`
	LastName  string    `json:"LastName" binding:"required"`
	Bio       string    `json:"Bio" binding:"-"`
	Birthday  time.Time `json:"Birthday" binding:"required"`
}

type LoginCredential struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
