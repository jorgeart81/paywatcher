package userdomain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Username string
	Password string
}

func (u *User) NewUser() *User {
	return &User{
		ID:       uuid.New(),
		Name:     u.Name,
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
	}
}
