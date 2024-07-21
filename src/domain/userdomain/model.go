package userdomain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Username string
	Password string
}

func NewUser(name, email, username, password string) *User {
	return &User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Username: username,
		Password: password,
	}
}
