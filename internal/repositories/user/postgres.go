package user

import (
	"context"
	"errors"
	"github.com/h-varmazyar/p3o/internal/entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const tableName = "users"

type Repository struct {
	*gorm.DB
	log *log.Logger
}

func New(log *log.Logger, db *gorm.DB) Repository {
	repo := Repository{
		DB:  db,
		log: log,
	}
	return repo
}

func (r Repository) Create(ctx context.Context, user entities.User) error {
	err := r.DB.WithContext(ctx).Save(&user).Error
	if err != nil {
		return ErrFailedToCreateUser.AddOriginalError(err)
	}
	return nil
}

func (r Repository) ReturnByMobile(ctx context.Context, mobile string) (entities.User, error) {
	user := entities.User{}
	err := r.DB.WithContext(ctx).Model(new(entities.User)).Where("mobile = ?", mobile).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrUserNotFound
		}
		return user, ErrFailedToFetchUser.AddOriginalError(err)
	}
	return user, nil
}

func (r Repository) ReturnByEmail(ctx context.Context, email string) (entities.User, error) {
	user := entities.User{}
	err := r.DB.WithContext(ctx).Model(new(entities.User)).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrUserNotFound
		}
		return user, ErrFailedToFetchUser.AddOriginalError(err)
	}
	return user, nil
}
