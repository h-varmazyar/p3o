package entities

import (
	"database/sql"
	"gorm.io/gorm"
)

type LinkStatus string

const (
	LinkStatusActive             LinkStatus = "ACTIVE"
	LinkStatusDeactivatedByUser  LinkStatus = "DEACTIVATED_BY_USER"
	LinkStatusDeactivatedByAdmin LinkStatus = "DEACTIVATED_BY_ADMIN"
)

var (
	linkStatusShowable = map[LinkStatus]string{
		LinkStatusActive:             "فعال",
		LinkStatusDeactivatedByUser:  "غیرفعال",
		LinkStatusDeactivatedByAdmin: "محدود شده",
	}
)

func (ls LinkStatus) ToShowableString() string {
	return linkStatusShowable[ls]
}

type Link struct {
	gorm.Model
	OwnerId    uint
	Key        string `gorm:"index:idx_link_key,unique"`
	RealLink   string
	Status     LinkStatus
	TotalVisit int64
	ExpireAt   sql.NullTime
	Immediate  bool
	Password   string
	MaxVisit   uint
}
