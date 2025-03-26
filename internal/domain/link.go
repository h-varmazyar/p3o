package domain

import "time"

type LinkCreateReq struct {
	UserId  uint   `json:"user_id"`
	Key     string `json:"key"`
	RealUrl string `json:"real_url"`
}

type LinkCreateResp struct {
	Key       string `json:"key"`
	Url       string `json:"url"`
	Status    string `json:"status"`
	Immediate bool   `json:"immediate"`
}

type Link struct {
	ID        uint      `json:"-"`
	ShortLink string    `json:"short_link"`
	Url       string    `json:"url"`
	Visits    uint      `json:"visits"`
	CreatedAt time.Time `json:"created_at"`
}

type All struct {
	Links []Link `json:"links"`
}
