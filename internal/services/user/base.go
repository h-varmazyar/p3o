package user

import (
	"context"
	sysErr "errors"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	ReturnByMobile(ctx context.Context, username string) (entities.User, error)
	ReturnByEmail(ctx context.Context, username string) (entities.User, error)
	Create(ctx context.Context, user entities.User) error
}

var signedKey = []byte(`f0dde8d96b9a581f24b1ae25f96cdd680f3bc50679274b18a4aa2a06d0c748b4b68b4d70ae991813f163cbf8ee64d89c63aa530793caaaed6580e1f9ad3bb6edff025ddaee4708a83dad233596e217aa40042055c13ee6a6d27fa942689d3fed6f1e07b26515087894698461797b831def9483144a1a49118f831e6fe439cf9779f8fb500e9ba562e81324bce752f1fb285f74ebfdc1ade78dbd90860abbcef2ca5ee238e7c0b141baa6f0fd2f26401c0d44be4a4a6ac086ac508e5451603a4f06b6f45d41fc53ae40c593c1cc4ddd2b56fc196446eeb944518ff7e946141700c747bc1c28b01d8d7398089fcc7a68bb5c42666ef14bb78c10e6d5194fbe4a383083e3a9d443ffd3d22b0c94bc44248ea8f4c8f4522f936f4435f068e0272825d191754d2f8e9036eb64ec21e7348216cae203ab07e060a9cc1b6ab9d1ebe34f293b548dcc978252ae22c6aad6f283b0ce4d2414ae20172e77e5eb22f16f3a495851ba458feadd561cf052a5d3136b750b01453f2adfab0d7b5154f8fde3f0ce4f848bd67bc72d6213db545fecaedd815ae26cb1b1f7404a8c64b198752d57a72127dbb01731bacf3be92131c171b8cbd208ae46d5c36fb6f9421a153c438693a3b0f4c2639aeaaecc5d9fe57390c3d82595ee4cb182cff604f8b408e40974a62f6e68f39f348d8f3ec19190164db61e3ed689c4de30188535705ea7569bd57e`)

type Service struct {
	log      *log.Logger
	userRepo UserRepository
	//signKey   *rsa.PrivateKey
	//verifyKey *rsa.PublicKey
	cfg configs.UserService
}

func New(log *log.Logger, cfg configs.UserService, userRepo UserRepository) (Service, error) {
	srv := Service{
		log:      log,
		userRepo: userRepo,
		cfg:      cfg,
	}
	//
	//if err := srv.generateKeys(); err != nil {
	//	return Service{}, err
	//}
	return srv, nil
}

//func (s Service) generateKeys() error {
//	var err error
//	s.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(s.cfg.JWTPublicKey))
//	if err != nil {
//		return err
//	}
//	s.signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(s.cfg.JWTPrivateKey))
//	if err != nil {
//		return err
//	}
//	return nil
//}

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
