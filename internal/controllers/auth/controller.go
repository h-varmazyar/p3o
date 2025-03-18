package auth

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"github.com/h-varmazyar/p3o/internal/controllers/user"
	"github.com/h-varmazyar/p3o/internal/repositories/auth"
	"go.uber.org/fx"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

type Configs struct {
	JWTPublicKey  string
	JWTPrivateKey string
}

type Controller struct {
	cfg         Configs
	userService user.Service
}

type Params struct {
	fx.In

	Cfg       *Configs
	UserModel auth.Model
}

type Result struct {
	fx.Out

	Controller *Controller
}

func New(p Params) Result {
	controller := &Controller{
		cfg:       *p.Cfg,
		userModel: p.UserModel,
	}
	if err := controller.generateKeys(); err != nil {
		panic(err)
	}
	return Result{Controller: controller}
}

func (c *Controller) generateKeys() error {
	var err error
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(c.cfg.JWTPublicKey))
	if err != nil {
		return err
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(c.cfg.JWTPrivateKey))
	if err != nil {
		return err
	}
	return nil
}
