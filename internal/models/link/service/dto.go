package service

type CreateLinkReq struct {
	Key     string
	RealUrl string
}

type FetchLinkReq struct {
	Key string
}

type Link struct {
	Url       string
	Immediate bool
}
