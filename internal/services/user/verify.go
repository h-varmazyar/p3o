package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (s Service) Verify(ctx context.Context, userId uint) (domain.VerifyUserResp, error) {
	_, err := s.userRepo.ReturnById(ctx, userId)
	if err != nil {
		return domain.VerifyUserResp{}, err
	}

	otpCode := utils.RandomOTP(5)

	//todo: send random via message

	if err = s.verificationCodeCache.Set(userId, otpCode); err != nil {
		return domain.VerifyUserResp{}, err
	}
	return domain.VerifyUserResp{}, nil
}
