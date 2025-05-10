package user

import (
	"context"
	sysErr "errors"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/h-varmazyar/p3o/pkg/cache"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	ReturnById(ctx context.Context, id uint) (entities.User, error)
	ReturnByMobile(ctx context.Context, username string) (entities.User, error)
	ReturnByEmail(ctx context.Context, username string) (entities.User, error)
	Create(ctx context.Context, user entities.User) (entities.User, error)
	Update(ctx context.Context, user entities.User) error
}

type Service struct {
	log                   *log.Logger
	cfg                   configs.UserService
	userRepo              UserRepository
	verificationCodeCache *cache.VerificationCodeRedisCache
}

func New(log *log.Logger, cfg configs.UserService, userRepo UserRepository, verificationCodeCache *cache.VerificationCodeRedisCache) (Service, error) {
	srv := Service{
		log:                   log,
		cfg:                   cfg,
		userRepo:              userRepo,
		verificationCodeCache: verificationCodeCache,
	}

	return srv, nil
}

func (s Service) fetchUser(ctx context.Context, username string) (entities.User, bool, error) {
	var err error
	user := entities.User{}
	if utils.IsValidMobile(username) {
		user, err = s.userRepo.ReturnByMobile(ctx, username)
		if err != nil {
			if sysErr.As(err, &errors.ErrUserNotFound) {
				user.Mobile = username
				return user, false, nil
			}
			return user, false, err
		}
	} else if utils.IsValidEmail(username) {
		user, err = s.userRepo.ReturnByEmail(ctx, username)
		if err != nil {
			if sysErr.As(err, &errors.ErrUserNotFound) {
				user.Email = &username
				return user, false, nil
			}
			return user, false, err
		}
	} else {
		return user, false, errors.ErrInvalidUsernamePassword
	}
	return user, true, nil
}
