package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/models/auth"
	"github.com/h-varmazyar/p3o/pkg/environments"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"time"
)

var (
	configs   *config
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func init() {
	configs = new(config)
	err := environments.LoadEnvironments(configs)
	if err != nil {
		panic(fmt.Sprintf("failed to load auth controller configs: %v", err))
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(configs.JWTPublicKey))
	if err != nil {
		panic(fmt.Sprintf("failed to parse public rsa key: %v", err))
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(configs.JWTPrivateKey))
	if err != nil {
		panic(fmt.Sprintf("failed to parse private rsa key: %v", err))
	}
}

type config struct {
	JWTPublicKey  string `env:"JWT_PUBLIC_KEY,file,required"`
	JWTPrivateKey string `env:"JWT_PRIVATE_KEY,file,required"`
}

type Controller struct {
	userModel auth.Model
}

type Params struct {
	fx.In

	UserModel auth.Model
}

type Result struct {
	fx.Out

	Controller *Controller
}

func New(p Params) Result {
	controller := &Controller{
		userModel: p.UserModel,
	}
	return Result{Controller: controller}
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
		if err = utils.CompareHashPassword(loginReq.Password, user.HashedPassword); err != nil {
			log.WithError(err).Error("failed to generage hashed password")
			utils.JsonHttpResponse(ctx, nil, ErrInvalidUsernamePassword.AddOriginalError(err), false)
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

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err := token.SignedString(signKey)
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
