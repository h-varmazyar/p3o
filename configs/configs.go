package configs

import (
	gormext "github.com/h-varmazyar/gopack/gorm"
	"github.com/h-varmazyar/p3o/internal/controllers"
	"github.com/h-varmazyar/p3o/pkg/db/redis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Configs struct {
	Controller  controllers.Configs `yaml:"controller"`
	GormConfigs gormext.Configs     `yaml:"gormConfigs"`
	RedisConfig redis.Configs       `yaml:"redisConfigs"`
}

type Params struct {
	fx.In

	Log *log.Logger
}

type Result struct {
	fx.Out

	Configs *Configs
}

func LoadConfigs(p Params) (Result, error) {
	p.Log.Infof("reding configs...")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		p.Log.Warnf("failed to read from env: %v", err)
		viper.AddConfigPath("./configs")  //path for docker compose configs
		viper.AddConfigPath("../configs") //path for local configs
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		if err = viper.ReadInConfig(); err != nil {
			p.Log.Errorf("failed to read configs")
			return Result{}, err
		}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		p.Log.Errorf("faeiled unmarshal")
		return Result{}, err
	}

	return Result{Configs: conf}, nil
}
