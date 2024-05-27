package entities

import (
	"gorm.io/gorm"
	"time"
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
	VerifiedAt     *time.Time
}

func (r UserRole) ToString() string {
	return string(r)
}
