package service

import (
	"context"
)

type Service interface {
	CreateLink(ctx context.Context, req *CreateLinkReq) (*Link, error)
	FetchLink(ctx context.Context, req *FetchLinkReq) (*Link, error)
}
