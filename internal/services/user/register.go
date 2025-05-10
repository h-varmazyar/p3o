package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/cache"
	"github.com/h-varmazyar/p3o/pkg/utils"
	"time"
)

const expirationDuration = time.Minute * 2

func (s Service) Register(ctx context.Context, req domain.RegisterUserReq) (domain.RegisterResp, error) {
	_, err := s.userRepo.ReturnByMobile(ctx, req.Mobile)
	if err == nil {
		return domain.RegisterResp{}, errors.ErrUserMobileAvailable
	}

	otpCode := utils.RandomOTP(6)

	//todo: send random via message

	vc := cache.VerificationCode{
		Code:     otpCode,
		Mobile:   req.Mobile,
		Password: req.Password,
	}
	if err = s.verificationCodeCache.Set(req.Mobile, vc, expirationDuration); err != nil {
		return domain.RegisterResp{}, err
	}

	resp := domain.RegisterResp{
		ExpirationDuration: expirationDuration / time.Second,
	}
	return resp, nil
}
