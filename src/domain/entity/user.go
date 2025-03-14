package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type UserEnt struct {
	ID              uuid.UUID  `json:"id"`
	Email           string     `json:"email"`
	NormalizedEmail string     `json:"normalizedEmail"`
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	Role            []string   `json:"role"`
	Active          bool       `json:"active"`
	SecurityStamp   string     `gorm:"column:securityStamp"`
	DeletedAt       *time.Time `json:"deletedAt"`
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
	u.NormalizedEmail = normalize(u.Email)
	return u
}

func normalize(email string) string {
	caser := cases.Upper(language.Und) // "Und" = sin especificar idioma
	normalizedEmail := caser.String(email)
	return strings.TrimSpace(normalizedEmail)
}
