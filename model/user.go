package model

import "time"

type SignupCredential struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Bio       string
	Birthday  time.Time
}

type LoginCredential struct {
	Username string
	Password string
}
