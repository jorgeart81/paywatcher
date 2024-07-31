package entity

import "github.com/google/uuid"

type UserEnt struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     []string  `json:"role"`
	Active   bool      `json:"active"`
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
		Active:   true,
	}
}

func (u *UserEnt) UpdateUser() *UserEnt {
	return &UserEnt{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
		Active:   u.Active,
	}
}
