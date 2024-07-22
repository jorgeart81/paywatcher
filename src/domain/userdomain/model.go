package userdomain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Username string
	Password string
	Role     []Role
}
type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

var AllRoles = []Role{AdminRole, UserRole}

func (u *User) NewUser() *User {
	return &User{
		ID:       uuid.New(),
		Name:     u.Name,
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     []Role{UserRole},
	}
}
