package auth

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
<<<<<<< HEAD
	"github.com/h-varmazyar/p3o/internal/models/auth"
=======
<<<<<<< HEAD
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/models/user"
	"github.com/h-varmazyar/p3o/pkg/environments"
	"github.com/h-varmazyar/p3o/pkg/utils"
	log "github.com/sirupsen/logrus"
=======
	"github.com/h-varmazyar/p3o/internal/models/auth"
>>>>>>> 3ffa25e (change project structure to service base)
>>>>>>> 292128d (feat: add link creation)
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
<<<<<<< HEAD
=======
<<<<<<< HEAD
	userModel user.Model
=======
>>>>>>> 292128d (feat: add link creation)
	cfg         Configs
	userService user.Service

	userModel auth.Model
<<<<<<< HEAD
=======
>>>>>>> 3ffa25e (change project structure to service base)
>>>>>>> 292128d (feat: add link creation)
}

type Params struct {
	fx.In

<<<<<<< HEAD
	Cfg       *Configs
	UserModel auth.Model
=======
<<<<<<< HEAD
	UserModel user.Model
=======
	Cfg       *Configs
	UserModel auth.Model
>>>>>>> 3ffa25e (change project structure to service base)
>>>>>>> 292128d (feat: add link creation)
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
