package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers"
	"github.com/h-varmazyar/p3o/internal/models/link"
	db2 "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log := new(logrus.Logger)
	ctx := context.Background()

	configs, err := loadConfigs(log)
	if err != nil {
		log.WithError(err).Panicf("failed to load configs")
	}

	db, err := db2.NewDatabase(ctx, *configs.DB)
	if err != nil {
		log.WithError(err).Panicf("failed to create database")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:       configs.RedisAddress,
		ClientName: "P3O",
		Username:   configs.RedisPassword,
		Password:   configs.RedisPassword,
		DB:         configs.LinkCacheDB,
	})

	linkApp, err := link.NewApp(ctx, log, db, redisClient, configs.LinkApp)
	if err != nil {
		log.WithError(err).Panicf("failed to create link app")
	}

	if err = controllers.NewController(log, linkApp.GetService()); err != nil {
		log.WithError(err).Panic("failed to start controllers")
	}

}

func loadConfigs(log *logrus.Logger) (*Configs, error) {
	log.Infof("reding configs...")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("failed to read from env: %v", err)
		viper.AddConfigPath("./configs")  //path for docker compose configs
		viper.AddConfigPath("../configs") //path for local configs
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		if err = viper.ReadInConfig(); err != nil {
			log.Errorf("failed to read configs")
			return nil, err
		}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		log.Errorf("faeiled unmarshal")
		return nil, err
	}

	return conf, nil
}
