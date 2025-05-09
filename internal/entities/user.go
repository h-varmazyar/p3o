package entities

import (
	"database/sql"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin = "admin"
	RoleUser  = "auth"
)

type User struct {
	gorm.Model
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	Email          string   `gorm:"unique" json:"email"`
	Mobile         string   `gorm:"unique" json:"mobile"`
	HashedPassword string   `json:"-"`
	Role           UserRole `json:"role"`
	VerifiedAt     sql.NullTime
}

func (r UserRole) ToString() string {
	return string(r)
}
