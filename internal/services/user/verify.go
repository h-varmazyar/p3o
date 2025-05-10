package user

import (
	"context"
	"database/sql"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/jwt"
	"github.com/h-varmazyar/p3o/pkg/utils"
	"time"
)

func (s Service) Verify(ctx context.Context, req domain.VerifyUserReq) (domain.Tokens, error) {
	vc, err := s.verificationCodeCache.Get(req.Mobile)
	if err != nil {
		return domain.Tokens{}, err
	}

	if vc.Code != req.Code {
		return domain.Tokens{}, errors.ErrWrongVerificationCode
	}

	if vc.Mobile != req.Mobile {
		return domain.Tokens{}, errors.ErrWrongVerificationCode
	}

	_, err = s.userRepo.ReturnByMobile(ctx, req.Mobile)
	if err == nil {
		return domain.Tokens{}, errors.ErrUserMobileAvailable
	}

	hashedPassword, err := utils.GenerateHashPassword(vc.Password)
	if err != nil {
		return domain.Tokens{}, err
	}

	user := entities.User{
		Mobile:         req.Mobile,
		HashedPassword: hashedPassword,
		Role:           entities.RoleUser,
	}
	user.VerifiedAt = sql.NullTime{Time: time.Now(), Valid: true}

	if user, err = s.userRepo.Create(ctx, user); err != nil {
		return domain.Tokens{}, err
	}

	jwtToken := jwt.GenerateToken(user.ID)

	return domain.Tokens{
		Token:      jwtToken.AccessToken,
		ExpireAt:   jwtToken.ExpiresAt,
		IsVerified: user.VerifiedAt.Valid,
	}, nil
}
