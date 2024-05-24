package auth

import (
	"context"
	"errors"
	"github.com/h-varmazyar/p3o/internal/entities"
	userRepository "github.com/h-varmazyar/p3o/internal/models/user"
	"github.com/h-varmazyar/p3o/pkg/utils"
)

func (c *Controller) fetchUser(ctx context.Context, username string) (*entities.User, bool, error) {
	var err error
	user := new(entities.User)
	if utils.IsValidMobile(username) {
		user, err = c.repository.ReturnByMobile(ctx, username)
		if err != nil {
			if errors.Is(err, userRepository.ErrUserNotFound) {
				user = &entities.User{Mobile: username}
				return user, false, nil
			}
			return nil, false, err
		}
	} else if utils.IsValidEmail(username) {
		user, err = c.repository.ReturnByEmail(ctx, username)
		if err != nil {
			if errors.Is(err, userRepository.ErrUserNotFound) {
				user = &entities.User{Email: username}
				return user, false, nil
			}
			return nil, false, err
		}
	} else {
		return nil, false, ErrInvalidUsernamePassword
	}
	return user, true, nil
}
