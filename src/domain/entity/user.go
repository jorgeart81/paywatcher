package entity

import "github.com/google/uuid"

type UserEnt struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     []string  `json:"role"`
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

var UserAllowedRoles = map[string]bool{
	RoleAdmin: true,
	RoleUser:  true,
}

func (u *UserEnt) NewUser() *UserEnt {
	return &UserEnt{
		ID:       uuid.New(),
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     []string{RoleUser},
	}
}

func (u *UserEnt) UpdateUser() *UserEnt {
	return &UserEnt{
		ID:       uuid.New(),
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     []string{RoleUser},
	}
}
