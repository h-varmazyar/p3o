package entities

import (
	"database/sql"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Key         string `gorm:"index:idx_link_key,unique"`
	RealLink    string
	UniqueVisit int64
	TotalVisit  int64
	ExpireAt    sql.NullTime
	Immediate   bool
}
