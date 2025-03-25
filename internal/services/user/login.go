package user

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

func (s Service) Login(ctx context.Context, username, password string) (domain.Tokens, error) {
	user, found, err := s.fetchUser(ctx, username)
	if err != nil {
		return domain.Tokens{}, err
	}

	if found {
		if err = utils.CompareHashPassword(password, user.HashedPassword); err != nil {
			log.WithError(err).Error("failed to generate hashed password")
			return domain.Tokens{}, errors.ErrWrongPassword.AddOriginalError(err)
		}
	} else {
		user.HashedPassword, err = utils.GenerateHashPassword(password)
		if err != nil {
			return domain.Tokens{}, errors.ErrPasswordHashingFailed.AddOriginalError(err)
		}
		err = s.userRepo.Create(ctx, user)
		if err != nil {
			return domain.Tokens{}, err
		}
	}

	expirationTime := time.Now().Add(30 * 124 * time.Hour)

	claims := &entities.Claims{
		Role: user.Role.ToString(),
		StandardClaims: jwt.StandardClaims{
			Subject:   username,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(signedKey)
	if err != nil {
		return domain.Tokens{}, errors.ErrLoginFailed.AddOriginalError(err)
	}

	return domain.Tokens{
		Token:      tokenString,
		ExpireAt:   expirationTime,
		IsVerified: user.VerifiedAt != nil,
	}, nil
}
