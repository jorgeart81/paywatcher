package userdomain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `form:"email" json:"email" binding:"email,required"`
	Username string    `form:"usename" json:"username"`
	Password string    `form:"password" json:"password,omitempty" binding:"required"`
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
