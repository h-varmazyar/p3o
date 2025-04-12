package domain

import "time"

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

type GetLinkReq struct {
	Key       string `json:"key"`
	OS 		  string `json:"os"`
	Browser   string `json:"browser"`
	Cookie    string `json:"coockie"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

type GetLinkResp struct {
	Url    string `json:"url"`
	Cookie string `json:"cookie"`
}

type Link struct {
	ShortLink string    `json:"short_link"`
	Url       string    `json:"url"`
	Visits    uint      `json:"visits"`
	CreatedAt time.Time `json:"created_at"`
}

type LinkList struct {
	Links []Link `json:"links"`
}

type ChartItem struct {
	Count     uint   `json:"count"`
	TimeLabel string `json:"time_label"`
}
