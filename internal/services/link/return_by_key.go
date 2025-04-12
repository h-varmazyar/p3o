package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/oklog/ulid/v2"
	"time"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) ReturnByKey(ctx context.Context, req domain.GetLinkReq) (domain.GetLinkResp, error) {
	link, err := s.linkRepo.ReturnByKey(ctx, req.Key)
	if err != nil {
		return domain.GetLinkResp{}, err
	}

	visit := entities.Visit{
		ID:        ulid.Make().String(),
		LinkId:    link.ID,
		UserId:    link.OwnerId,
		OS:        req.OS,
		Browser:   req.Browser,
		UserAgent: req.UserAgent,
		IP:        req.IP,
	}

	if req.Cookie == "" {
		visit.Cookie = createCookie(visit.ID, req.Key)
	}

	url := ""
	if link.Immediate {
		url = link.RealLink
		visit.RedirectedAt = time.Now()
		visit.Status = entities.VisitStatusCompleted
	} else {
		url = fmt.Sprintf("%v/%v/%v", s.cfg.IndirectBaseURL, link.Key, visit.ID)
		visit.Status = entities.VisitStatusAdsPending
	}

	s.visitChan <- visit

	return domain.GetLinkResp{
		Url: url,
	}, nil
}

func createCookie(id, key string) (string, error) {
	str:=id+key
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}