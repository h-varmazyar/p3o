package link

import (
	"context"
	"database/sql"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
	"strings"
	"time"
)

func (s Service) Edit(ctx context.Context, req domain.EditLinkReq) error {
	link, err := s.linkRepo.ReturnByKey(ctx, req.Key)
	if err != nil {
		return err
	}

	if link.OwnerId != req.UserId {
		return errors.ErrLinkOwnerMismatch
	}

	if ok, expireAt := canChangeExpireAt(req.ExpireAt, link.ExpireAt); ok {
		link.ExpireAt = expireAt
	}

	if ok, status := canChangeStatus(req.Status, link.Status); ok {
		link.Status = status
	}

	if ok, password := canChangePassword(req.Password, link.Password); ok {
		link.Password = password
	}

	if req.MaxVisit != link.MaxVisit {
		link.MaxVisit = req.MaxVisit
	}

	if req.Immediate != link.Immediate {
		link.Immediate = req.Immediate
	}

	if err = s.linkRepo.Update(ctx, link); err != nil {
		return err
	}

	return nil
}

func canChangeExpireAt(expireAt *time.Time, linkExpireAt sql.NullTime) (bool, sql.NullTime) {
	if linkExpireAt.Valid && (expireAt == nil || *expireAt != linkExpireAt.Time) {
		return true, sql.NullTime{
			Valid: false,
		}
	} else if !linkExpireAt.Valid && expireAt != nil {
		return true, sql.NullTime{
			Valid: true,
			Time:  *expireAt,
		}
	}

	return false, sql.NullTime{}
}

func canChangeStatus(status string, linkStatus entities.LinkStatus) (bool, entities.LinkStatus) {
	switch strings.ToLower(status) {
	case "active":
		if linkStatus == entities.LinkStatusDeactivatedByUser {
			return true, entities.LinkStatusActive
		}
	case "deactivate":
		if linkStatus == entities.LinkStatusActive {
			return true, entities.LinkStatusDeactivatedByUser
		}
	}

	return false, ""
}

func canChangePassword(password *string, linkHashedPassword string) (bool, string) {
	if password == nil {
		if linkHashedPassword == "" {
			return false, ""
		}
		return true, ""
	} else {
		if utils.CompareHashPassword(*password, linkHashedPassword) == nil {
			return false, ""
		}
		if hashedPassword, err := utils.GenerateHashPassword(*password); err == nil {
			return true, hashedPassword
		}
		return false, ""
	}
}
