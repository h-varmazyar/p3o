package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func main() {
	//log := new(logrus.Logger)
	//ctx := context.Background()
	//
	//configs, err := loadConfigs(log)
	//if err != nil {
	//	log.WithError(err).Panicf("failed to load configs")
	//}

	//db, err := db2.NewDatabase(ctx, *configs.DB)
	//if err != nil {
	//	log.WithError(err).Panicf("failed to create database")
	//}
	//
	//redisClient := redis.NewClient(&redis.Options{
	//	Addr:       configs.RedisAddress,
	//	ClientName: "P3O",
	//	Username:   configs.RedisPassword,
	//	Password:   configs.RedisPassword,
	//	DB:         configs.LinkCacheDB,
	//})
	//
	//linkApp, err := link.NewApp(ctx, log, db, redisClient, configs.LinkApp)
	//if err != nil {
	//	log.WithError(err).Panicf("failed to create link app")
	//}
	//
	//if err = controllers.NewController(log, linkApp.GetService()); err != nil {
	//	log.WithError(err).Panic("failed to start controllers")
	//}

	fx := initializeDependencies()

	fx.Run()

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

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8080"}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
