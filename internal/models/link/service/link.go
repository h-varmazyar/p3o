package service

import (
	"context"
	"errors"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/models/link/repository"
	"github.com/h-varmazyar/p3o/internal/models/link/workers"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/url"
)

type LinkService struct {
	log       *log.Logger
	persistDB repository.Repository
	cacheDB   repository.Repository
	visitChan chan workers.VisitRecord
}

func NewLinkService(log *log.Logger, persistDB, cacheDB repository.Repository, visitChannel chan workers.VisitRecord) Service {
	return &LinkService{
		log:       log,
		persistDB: persistDB,
		cacheDB:   cacheDB,
		visitChan: visitChannel,
	}
}

func (s *LinkService) CreateLink(ctx context.Context, req *CreateLinkReq) (*Link, error) {
	urlValue, err := url.Parse(req.RealUrl)
	if err != nil {
		return nil, err
	}

	if req.Key == "" {
		req.Key = generateKey(urlValue)
	}

	link := &entities.Link{
		Model:     gorm.Model{},
		Key:       req.Key,
		RealLink:  urlValue.String(),
		Immediate: true,
	}

	if err = s.persistDB.Create(ctx, link); err != nil {
		return nil, err
	}

	resp := &Link{
		Url:       link.RealLink,
		Immediate: link.Immediate,
	}

	return resp, nil
}

func (s *LinkService) FetchLink(ctx context.Context, req *FetchLinkReq) (*Link, error) {
	link, err := s.cacheDB.ReturnByKey(ctx, req.Key)
	if err != nil && errors.As(err, &repository.ErrCacheFetchFailed) {
		return nil, err
	}

	if link == nil {
		link, err = s.persistDB.ReturnByKey(ctx, req.Key)
		if err != nil {
			return nil, err
		}

		if err = s.cacheDB.Create(ctx, link); err != nil {
			return nil, err
		}
	}

	resp := new(Link)
	resp.Url = link.RealLink
	resp.Immediate = link.Immediate

	s.visitChan <- workers.VisitRecord{LinkId: link.ID}

	return resp, nil
}
