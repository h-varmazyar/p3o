package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (s Service) ChangePassword(ctx context.Context, req domain.ChangePasswordReq) error {
	user, err := s.userRepo.ReturnById(ctx, req.UserId)
	if err != nil {
		return err
	}

	if err = utils.CompareHashPassword(req.OldPassword, user.HashedPassword); err != nil {
		return err
	}

	user.HashedPassword, err = utils.GenerateHashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil

}
