package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserEnt struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Role      []string   `json:"role"`
	Active    bool       `json:"active"`
	DeletedAt *time.Time `json:"deletedAt"`
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
	u.ID = uuid.New()
	u.Role = []string{RoleUser}
	u.Active = true
	return u
}
