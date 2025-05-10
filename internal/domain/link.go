package domain

import (
	"github.com/h-varmazyar/p3o/pkg/pagination"
	"time"
)

type LinkCreateReq struct {
	UserId    uint   `json:"user_id"`
	Key       string `json:"key"`
	RealUrl   string `json:"real_url"`
	Immediate bool   `json:"immediate"`
}

type LinkCreateResp struct {
	Key       string `json:"key"`
	Url       string `json:"url"`
	Status    string `json:"status"`
	Immediate bool   `json:"immediate"`
}

type Link struct {
	ShortLink   string    `json:"short_link"`
	Url         string    `json:"url"`
	Status      string    `json:"status"`
	IsImmediate bool      `json:"is_immediate"`
	Visits      uint      `json:"visits"`
	CreatedAt   time.Time `json:"created_at"`
}

type LinkDetails struct {
	ShortLink           string    `json:"short_link"`
	Url                 string    `json:"url"`
	Status              string    `json:"status"`
	IsImmediate         bool      `json:"is_immediate"`
	Visits              uint      `json:"visits"`
	ExpireAt            time.Time `json:"expired_at"`
	CreatedAt           time.Time `json:"created_at"`
	ProtectedByPassword bool      `json:"protected_by_password"`
	MaxVisit            uint32    `json:"max_visit"`
}

type LinkList struct {
	Links      []Link                `json:"links"`
	Pagination pagination.Pagination `json:"pagination"`
}

type ChartItem struct {
	Count     uint   `json:"count"`
	TimeLabel string `json:"time_label"`
}
