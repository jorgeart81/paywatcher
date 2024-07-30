package request

import (
	"fmt"
	"paywatcher/src/domain/entity"
	"regexp"
)

type RegisterUser struct {
	Email    string   `form:"email" json:"email" binding:"email,required"`
	Username string   `form:"usename" json:"username" binding:"required"`
	Password string   `form:"password" json:"password,omitempty" binding:"required"`
	Role     []string `form:"role" json:"role"`
}

func (u *RegisterUser) ToUserEntity() *entity.UserEnt {
	return &entity.UserEnt{
		Email:    u.Email,
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
}

// ValidateRoles checks if all roles assigned to the user are allowed.
func (u *RegisterUser) ValidateRoles() error {
	for _, role := range u.Role {
		if _, ok := entity.UserAllowedRoles[role]; !ok {
			return fmt.Errorf("invalid role: %s", role)
		}
	}
	return nil
}

func (u *RegisterUser) ValidatePassword() error {
	password := u.Password
	//  "password must be at least 8 characters long, include one uppercase letter, one lowercase letter, one number, and one special character"
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	upper := regexp.MustCompile(`[A-Z]`)
	if !upper.MatchString(password) {
		return fmt.Errorf("password must include at least one uppercase letter")
	}

	lower := regexp.MustCompile(`[a-z]`)
	if !lower.MatchString(password) {
		return fmt.Errorf("password must include at least one lowercase letter")
	}

	number := regexp.MustCompile(`\d`)
	if !number.MatchString(password) {
		return fmt.Errorf("password must include at least one number")
	}

	special := regexp.MustCompile(`[@$!%*?&]`)
	if !special.MatchString(password) {
		return fmt.Errorf("password must include at least one special character %s", `(@$!%*?&)`)
	}
	return nil
}

type LoginUser struct {
	Email    string `form:"email" json:"email" binding:"email,required"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
}
