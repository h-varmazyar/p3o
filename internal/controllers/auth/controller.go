package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/h-varmazyar/p3o/internal/entities"
	user "github.com/h-varmazyar/p3o/internal/models/auth"
	"github.com/h-varmazyar/p3o/pkg/utils"
	"go.uber.org/fx"
	"time"
)

type Controller struct {
	userModel user.Model
	configs   *configs.ControllerConfigs
}

func New(lc fx.Lifecycle, configs *configs.ControllerConfigs, userModel user.Model) *Controller {
	controller := &Controller{
		userModel: userModel,
		configs:   configs,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return controller
}

func (c *Controller) Login(ctx *gin.Context) {
	loginReq := new(LoginReq)

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		utils.JsonHttpResponse(ctx, nil, ErrLoginFailed.AddOriginalError(err), false)
		return
	}

	user, found, err := c.fetchUser(ctx, loginReq.Username)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, err, false)
		return
	}

	if found {
		if !utils.CompareHashPassword(loginReq.Password, user.HashedPassword) {
			utils.JsonHttpResponse(ctx, nil, ErrInvalidUsernamePassword, false)
			return
		}
	} else {
		user.HashedPassword, err = utils.GenerateHashPassword(loginReq.Password)
		if err != nil {
			utils.JsonHttpResponse(ctx, nil, ErrPasswordHashingFailed.AddOriginalError(err), false)
			return
		}
		err = c.userModel.Create(ctx, user)
		if err != nil {
			utils.JsonHttpResponse(ctx, nil, err, false)
			return
		}
	}

	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	claims := &entities.Claims{
		Role: user.Role.ToString(),
		StandardClaims: jwt.StandardClaims{
			Subject:   loginReq.Username,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(c.configs.JWTSecret)
	if err != nil {
		utils.JsonHttpResponse(ctx, nil, ErrLoginFailed.AddOriginalError(err), false)
		return
	}

	resp := &LoginResp{
		Token:        tokenString,
		ExpireAt:     expirationTime,
		VerifiedUser: user.VerifiedAt != nil,
	}

	utils.JsonHttpResponse(ctx, resp, nil, true)
}

func (c *Controller) Verify(ctx *gin.Context) {

}

func (c *Controller) Logout(ctx *gin.Context) {

}
