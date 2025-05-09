package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
)

func (s Service) EditUser(ctx context.Context, req domain.EditUserReq) error {
	user, err := s.userRepo.ReturnById(ctx, req.UserId)
	if err != nil {
		return err
	}

	user.LastName = req.LastName
	user.FirstName = req.FirstName
	user.Email = req.Email

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
