package auth

import (
	"context"
	"errors"
	"github.com/h-varmazyar/p3o/internal/entities"
	userRepository "github.com/h-varmazyar/p3o/internal/models/user"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func (c *Controller) fetchUser(ctx context.Context, username string) (*entities.User, bool, error) {
	var err error
	user := new(entities.User)
	if utils.IsValidMobile(username) {
		user, err = c.userModel.ReturnByMobile(ctx, username)
		if err != nil {
			if errors.Is(err, userRepository.ErrUserNotFound) {
				user = &entities.User{Mobile: username}
				return user, false, nil
			}
			return nil, false, err
		}
	} else if utils.IsValidEmail(username) {
		user, err = c.userModel.ReturnByEmail(ctx, username)
		if err != nil {
			if errors.Is(err, userRepository.ErrUserNotFound) {
				user = &entities.User{Email: username}
				return user, false, nil
			}
			return nil, false, err
		}
	} else {
		log.Infof("username: %v", username)
		return nil, false, ErrInvalidUsernamePassword
	}
	return user, true, nil
}
