package user

import (
	"context"
	sysErr "errors"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/cache"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (s Service) Verify(ctx context.Context, req domain.VerifyUserReq) (domain.VerifyUserResp, error) {
	user, err := s.userRepo.ReturnById(ctx, req.UserId)
	if err != nil {
		return domain.VerifyUserResp{}, err
	}

	tempUser, err := s.userRepo.ReturnByMobile(ctx, req.Mobile)
	if err == nil || !sysErr.As(err, &errors.ErrUserNotFound) {
		return domain.VerifyUserResp{}, errors.ErrUserMobileAvailable
	}
	if tempUser.ID != user.ID {
		return domain.VerifyUserResp{}, errors.ErrUserMobileAvailable
	}

	otpCode := utils.RandomOTP(5)

	//todo: send random via message

	vc := cache.VerificationCode{
		Code:   otpCode,
		Mobile: req.Mobile,
	}
	if err = s.verificationCodeCache.Set(req.UserId, vc); err != nil {
		return domain.VerifyUserResp{}, err
	}
	return domain.VerifyUserResp{}, nil
}
