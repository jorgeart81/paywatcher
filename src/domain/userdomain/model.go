package userdomain

import "github.com/google/uuid"

type User struct {
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

var AllowedRoles = map[string]bool{
	RoleAdmin: true,
	RoleUser:  true,
}

func (u *User) NewUser() *User {
	return &User{
		ID:       uuid.New(),
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     []string{RoleUser},
	}
}
