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

	fields := make(map[string]any)

	if ok, expireAt := canChangeExpireAt(req.ExpireAt, link.ExpireAt); ok {
		fields["expire_at"] = expireAt
	}

	if ok, status := canChangeStatus(req.Status, link.Status); ok {
		fields["status"] = status
	}

	if ok, password := canChangePassword(req.Password, link.Password); ok {
		fields["password"] = password
	}

	if ok, maxVisit := canChangeMaxVisit(req.MaxVisit, link.MaxVisit); ok {
		fields["max_visit"] = maxVisit
	}

	if ok, immediate := canChangeImmediate(req.Immediate, link.Immediate); ok {
		fields["immediate"] = immediate
	}

	if len(fields) > 0 {
		if err = s.linkRepo.UpdateFields(ctx, link.ID, fields); err != nil {
			return err
		}
	}

	return nil
}

func canChangeExpireAt(expireAt *time.Time, linkExpireAt sql.NullTime) (bool, sql.NullTime) {
	if expireAt == nil {
		if linkExpireAt.Valid {
			return true, sql.NullTime{
				Valid: false,
			}
		}
	} else {
		if !linkExpireAt.Valid || *expireAt != linkExpireAt.Time {
			return true, sql.NullTime{
				Valid: true,
				Time:  *expireAt,
			}
		}
	}
	return false, sql.NullTime{}
}

func canChangeStatus(status *string, linkStatus entities.LinkStatus) (bool, entities.LinkStatus) {
	if status == nil {
		return false, linkStatus
	}
	switch strings.ToLower(*status) {
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

func canChangeMaxVisit(maxVisit *uint, linkMaxVisit uint) (bool, uint) {
	if maxVisit == nil {
		if linkMaxVisit > 0 {
			return true, 0
		}
		return false, 0
	}
	if *maxVisit != linkMaxVisit {
		return true, *maxVisit
	}
	return false, 0
}

func canChangeImmediate(immediate *bool, linkImmediate bool) (bool, bool) {
	if immediate == nil {
		return false, linkImmediate
	}
	if *immediate != linkImmediate {
		return true, *immediate
	}
	return false, false

}
