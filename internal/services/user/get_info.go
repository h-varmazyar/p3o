package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
)

func (s Service) Get(ctx context.Context, userId uint) (domain.UserInfo, error) {
	user, err := s.userRepo.ReturnById(ctx, userId)
	if err != nil {
		return domain.UserInfo{}, err
	}

	resp := domain.UserInfo{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Mobile:     user.Mobile,
		IsVerified: user.VerifiedAt.Valid,
	}

	if user.Email != nil {
		resp.Email = *user.Email
	}

	return resp, nil
}
