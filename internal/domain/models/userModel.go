package models

import "github.com/google/uuid"

type User struct {
	email    string
	password string
	id       uuid.UUID
	username string
}

func NewUser(username, email, password string) *User {
	return &User{
		email:    email,
		password: password,
		id:       uuid.New(),
		username: username,
	}
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetID() uuid.UUID {
	return u.id
}

func (u *User) ChangeEmail(email string) {
	u.email = email
}

func (u *User) ChangePassword(password string) {
	u.password = password
}
