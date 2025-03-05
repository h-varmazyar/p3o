package entities

import (
	"database/sql"
	"gorm.io/gorm"
)

type LinkStatus string

const(
	LinkStatusActive			 LinkStatus = "ACTIVE"
	LinkStatusDeactivatedByUser  LinkStatus = "DEACTIVATED_BY_USER"
	LinkStatusDeactivatedByAdmin LinkStatus = "DEACTIVATED_BY_ADMIN"
)

type Link struct {
	gorm.Model
	OwnerId 	uint
	Key         string `gorm:"index:idx_link_key,unique"`
	RealLink    string
	Status 		LinkStatus
	UniqueVisit int64
	TotalVisit  int64
	ExpireAt    sql.NullTime
	Immediate   bool
}
