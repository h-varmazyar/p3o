package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/configs"
	authController "github.com/h-varmazyar/p3o/internal/controllers/auth"
	linkController "github.com/h-varmazyar/p3o/internal/controllers/link"
	authModel "github.com/h-varmazyar/p3o/internal/models/auth"
	linkModel "github.com/h-varmazyar/p3o/internal/models/link"
	"github.com/h-varmazyar/p3o/internal/router"
	v1Router "github.com/h-varmazyar/p3o/internal/router/v1"
	"github.com/h-varmazyar/p3o/internal/workers"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var ctx = func() context.Context {
	return context.Background()
}

var generalDependenciesModule = fx.Module(
	"general",
	fx.Provide(
		ctx,
		log.New,
		db.NewDatabase,
		//initializeRedis,
		initializeGin,
		initializeVisitChannel,
	),
	fx.Invoke(func(log *log.Logger) {
		fx.WithLogger(log)
		log.Infof("redirecting fx logger to app logger")
	}),
	//fx.Invoke(func(redis *redis.Client, ctx context.Context) {
	//	ping, err := redis.Ping(ctx).Result()
	//	log.Infof("redis ping: %v - %v", ping, err)
	//}),
)

var modelsDependenciesModule = fx.Module(
	"models",
	fx.Provide(
		authModel.New,
		linkModel.New,
	),
)

var cacheDependenciesModule = fx.Module(
	"cache",
	fx.Provide(),
)

var workersDependenciesModule = fx.Module(
	"workers",
	fx.Provide(
		workers.NewVisitWorker,
	),
	fx.Invoke(func(log *log.Logger, visitWorker *workers.VisitsWorker) {
		log.Infof("Invoking visits worker")
	}),
)

var controllersDependenciesModule = fx.Module(
	"controllers",
	fx.Provide(
		authController.New,
		linkController.New,
	),
)

var routersDependenciesModule = fx.Module(
	"routers",
	fx.Provide(
		v1Router.New,
		router.New,
	),
	fx.Invoke(func(log *log.Logger, router *router.Router) {
		log.Infof("invoking router")
	}),
)

func initializeDependencies() *fx.App {
	diApp := fx.New(
		generalDependenciesModule,
		modelsDependenciesModule,
		controllersDependenciesModule,
		routersDependenciesModule,
		workersDependenciesModule,
		cacheDependenciesModule,
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

func initializeGin(log *log.Logger) *gin.Engine {
	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()
	g := gin.Default()
	return g
}

func initializeVisitChannel() chan workers.VisitRecord {
	return make(chan workers.VisitRecord, 10000)
}
