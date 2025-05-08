package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/jwt"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func (s Service) Login(ctx context.Context, username, password string) (domain.Tokens, error) {
	user, found, err := s.fetchUser(ctx, username)
	if err != nil {
		s.log.WithError(err)
		return domain.Tokens{}, err
	}

	s.log.Infof("user found: %v - %v", found, user)

	if found {
		if err = utils.CompareHashPassword(password, user.HashedPassword); err != nil {
			log.WithError(err).Error("failed to generate hashed password")
			return domain.Tokens{}, errors.ErrWrongPassword.AddOriginalError(err)
		}
	} else {
		s.log.Infof("creating new user")
		user.HashedPassword, err = utils.GenerateHashPassword(password)
		if err != nil {
			s.log.WithError(err).Error(errors.ErrPasswordHashingFailed.Code)
			return domain.Tokens{}, errors.ErrPasswordHashingFailed.AddOriginalError(err)
		}
		s.log.Infof("send new user to repo")
		err = s.userRepo.Create(ctx, user)
		if err != nil {
			s.log.WithError(err)
			return domain.Tokens{}, err
		}
		s.log.Infof("user creation done")
	}

	jwtToken := jwt.GenerateToken(user.ID)

	return domain.Tokens{
		Token:      jwtToken.AccessToken,
		ExpireAt:   jwtToken.ExpiresAt,
		IsVerified: user.VerifiedAt != nil,
	}, nil
}
