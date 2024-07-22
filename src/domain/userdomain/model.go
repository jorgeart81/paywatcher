package userdomain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     []Role    `json:"role"`
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
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     []Role{UserRole},
	}
}
