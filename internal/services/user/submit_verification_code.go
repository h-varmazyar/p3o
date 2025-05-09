package user

import (
	"context"
	"database/sql"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"time"
)

func (s Service) SubmitVerificationCode(ctx context.Context, req domain.SubmitVerificationCodeReq) error {
	user, err := s.userRepo.ReturnById(ctx, req.UserId)
	if err != nil {
		return err
	}

	vc, err := s.verificationCodeCache.Get(user.ID)
	if err != nil {
		return err
	}

	if vc.Code != req.Code {
		return errors.ErrWrongVerificationCode
	}

	if vc.Mobile != req.Mobile {
		return errors.ErrWrongVerificationCode
	}

	if vc.Mobile != req.Mobile {
		user.Mobile = vc.Mobile
	}
	user.VerifiedAt = sql.NullTime{Time: time.Now(), Valid: true}

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
