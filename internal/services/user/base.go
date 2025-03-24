package user

import (
	"context"
	"crypto/rsa"
	sysErr "errors"
	"github.com/golang-jwt/jwt"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

type userRepository interface {
	ReturnByMobile(ctx context.Context, username string) (entities.User, error)
	ReturnByEmail(ctx context.Context, username string) (entities.User, error)
	Create(ctx context.Context, user entities.User) error
}
type Service struct {
	userRepo  userRepository
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	cfg       Configs
}

func (s Service) generateKeys() error {
	var err error
	s.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(s.cfg.JWTPublicKey))
	if err != nil {
		return err
	}
	s.signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(s.cfg.JWTPrivateKey))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) fetchUser(ctx context.Context, username string) (entities.User, bool, error) {
	var err error
	user := entities.User{}
	if utils.IsValidMobile(username) {
		user, err = s.userRepo.ReturnByMobile(ctx, username)
		if err != nil {
			if sysErr.Is(err, errors.ErrUserNotFound) {
				user.Mobile = username
				return user, false, nil
			}
			return user, false, err
		}
	} else if utils.IsValidEmail(username) {
		user, err = s.userRepo.ReturnByEmail(ctx, username)
		if err != nil {
			if sysErr.Is(err, errors.ErrUserNotFound) {
				user.Email = username
				return user, false, nil
			}
			return user, false, err
		}
	} else {
		return user, false, errors.ErrInvalidUsernamePassword
	}
	return user, true, nil
}
