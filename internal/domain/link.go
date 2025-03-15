package domain

type LinkCreateReq struct {
	UserId uint    `json:"user_id"`
	Key     string `json:"key"`
	RealUrl string `json:"real_url"`
}

type LinkCreateResp struct {
	Key         string `json:"key"`
	Url    		string `json:"url"`
	Status 		string `json:"status"`
	Immediate   bool   `json:"immediate"`
}