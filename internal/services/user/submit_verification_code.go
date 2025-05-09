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

	if user.Mobile != req.Mobile {
		return errors.ErrMobileMismatch
	}

	code, err := s.verificationCodeCache.Get(user.ID)
	if err != nil {
		return err
	}

	if code != req.Code {
		return errors.ErrWrongVerificationCode
	}

	user.VerifiedAt = sql.NullTime{Time: time.Now(), Valid: true}

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
