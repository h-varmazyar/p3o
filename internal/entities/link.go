package entities

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	gorm.Model
	ShortCode   string `gorm:"index:idx_link_short_code,unique"`
	RealLink    string
	UniqueVisit int64
	TotalVisit  int64
	ExpireAt    time.Time
	Immediate   bool
}
