package configs

import (
	gormext "github.com/h-varmazyar/gopack/gorm"
	userService "github.com/h-varmazyar/p3o/internal/services/user"
	"github.com/h-varmazyar/p3o/pkg/db/redis"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Configs struct {
	GormConfigs gormext.Configs `yaml:"gormConfigs"`
	RedisConfig redis.Configs   `yaml:"redisConfigs"`

	UserServiceConfigs userService.Configs
}

type Params struct {
	fx.In

	Log *log.Logger
}

type Result struct {
	fx.Out

	Configs *Configs
}
