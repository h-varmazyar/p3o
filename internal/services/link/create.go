package link

import (
	"context"

	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) Create(ctx context.Context, req domain.LinkCreateReq) (domain.LinkCreateResp, error) {
	if !isValidLink(req.RealUrl) {
		return domain.LinkCreateResp{}, errors.ErrInvalidLink
	}

	if req.Key == "" {
		var err error
		req.Key, err = pickKey()
		if err != nil {
			return domain.LinkCreateResp{}, errors.ErrKeyGenerationFailed
		}
	}

	linkData := entities.Link{
		OwnerId:   req.UserId,
		Key:       req.Key,
		RealLink:  req.RealUrl,
		Immediate: req.Immediate,
		Status:    entities.LinkStatusActive,
	}

	link, err := s.linkRepo.Create(ctx, linkData)
	if err != nil {
		return domain.LinkCreateResp{}, errors.ErrUnexpected.AddOriginalError(err)
	}

	return domain.LinkCreateResp{
		Url:       linkData.RealLink,
		Key:       linkData.Key,
		Status:    string(link.Status),
		Immediate: linkData.Immediate,
	}, nil
}
