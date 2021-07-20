package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64
	Username       string
	hashedPassword []byte
}

func NewUser(id int64, username string, hashedPassword []byte) *User {
	return &User{
		ID:             id,
		Username:       username,
		hashedPassword: hashedPassword,
	}
}

func (u User) VerifyPassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(u.hashedPassword, password)
}
