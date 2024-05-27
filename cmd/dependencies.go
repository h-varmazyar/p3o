package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/h-varmazyar/p3o/internal/router"
	v1Router "github.com/h-varmazyar/p3o/internal/router/v1"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

var ctx = func() context.Context {
	return context.Background()
}

var generalDependenciesModule = fx.Module(
	"general",
	fx.Provide(
		ctx,
		log.New,
		configs.LoadConfigs,
		db.NewDatabase,
		initializeRedis,
	),
	fx.Invoke(func(log *log.Logger) {
		fx.Logger(log)
		log.Infof("redirecting fx logger to app logger")
	}),
	fx.Invoke(func(redis *redis.Client, ctx context.Context) {
		ping, err := redis.Ping(ctx).Result()
		log.Infof("redis ping: %v - %v", ping, err)
	}),
)

var modelsDependenciesModule = fx.Module(
	"models",
	fx.Provide(),
)

var controllersDependenciesModule = fx.Module(
	"controllers",
	fx.Provide(),
)

var routersDependenciesModule = fx.Module(
	"routers",
	fx.Provide(
		v1Router.New,
		router.New,
	),
)

func initializeDependencies() *fx.App {
	diApp := fx.New(
		generalDependenciesModule,
		modelsDependenciesModule,
		controllersDependenciesModule,
		routersDependenciesModule,
	)

	return diApp
}

func initializeRedis(configs *configs.Configs) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:       configs.RedisConfig.RedisAddress,
		ClientName: "P3O",
		Username:   configs.RedisConfig.RedisPassword,
		Password:   configs.RedisConfig.RedisPassword,
		DB:         configs.RedisConfig.LinkCacheDB,
	})
}

func initializeGin(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			s := &Handler{
				GinServer: r,
			}
			s.Say()
			return srv.ListenAndServe()
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return r
}

type Handler struct {
	GinServer *gin.Engine
}

func (s *Handler) Say() {
	s.GinServer.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
