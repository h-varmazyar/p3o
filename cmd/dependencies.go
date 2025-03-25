package main

import (
	"github.com/gin-gonic/gin"
	gormext "github.com/h-varmazyar/gopack/gorm"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/h-varmazyar/p3o/internal/workers"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//var ctx = func() context.Context {
//	return context.Background()
//}
//
//var generalDependenciesModule = fx.Module(
//	"general",
//	fx.Provide(
//		ctx,
//		log.New,
//		configs.New,
//		initializePostgres,
//		initializeRedis,
//		initializeGin,
//		initializeVisitChannel,
//	),
//	fx.Invoke(func(log *log.Logger) {
//		fx.WithLogger(log)
//		log.Infof("redirecting fx logger to app logger")
//	}),
//)
//
//var repositoriesDependenciesModule = fx.Module(
//	"repositories",
//	fx.Provide(
//		//userRepository.New,
//		linkRepository.New,
//		fx.Annotate(
//			userRepository.NewRepo,
//			fx.As(new(userService.UserRepository)),
//		),
//	),
//	fx.Invoke(func(log *log.Logger, visitWorker *workers.VisitsWorker) {
//		log.Infof("Invoking repositories")
//	}),
//)
//
//var cacheDependenciesModule = fx.Module(
//	"cache",
//	fx.Provide(),
//)
//
//var workersDependenciesModule = fx.Module(
//	"workers",
//	fx.Provide(
//		workers.NewVisitWorker,
//	),
//	fx.Invoke(func(log *log.Logger, visitWorker *workers.VisitsWorker) {
//		log.Infof("Invoking visits worker")
//	}),
//)
//
//var servicesDependenciesModule = fx.Module(
//	"services",
//	fx.Provide(
//		fx.Annotate(
//			userService.NewService,
//			fx.As(new(authController.UserService)),
//		),
//	),
//)
//
//var controllersDependenciesModule = fx.Module(
//	"controllers",
//	fx.Provide(
//		authController.New,
//		//linkController.New,
//	),
//)
//
//var routersDependenciesModule = fx.Module(
//	"routers",
//	fx.Provide(
//		v1Router.New,
//		router.New,
//	),
//	fx.Invoke(func(log *log.Logger, router *router.Router) {
//		log.Infof("invoking router")
//	}),
//)
//
//func initializeDependencies() *fx.App {
//	diApp := fx.New(
//		generalDependenciesModule,
//		repositoriesDependenciesModule,
//		servicesDependenciesModule,
//		controllersDependenciesModule,
//		routersDependenciesModule,
//		workersDependenciesModule,
//		cacheDependenciesModule,
//	)
//
//	return diApp
//}

func initializePostgres(configs gormext.Configs) (*gorm.DB, error) {
	configs.DbType = gormext.PostgreSQL
	db, err := gormext.Open(configs)
	if err != nil {
		log.WithError(err).Error("failed to open database")
		return nil, err
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = gormext.EnableExtensions(tx,
			gormext.UUIDExtension,
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Error("failed to add extensions to database")
		return nil, err
	}

	return db, nil
}

func initializeRedis(cfg configs.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:       cfg.Address,
		ClientName: "P3O",
		Username:   cfg.Username,
		Password:   cfg.Password,
		DB:         cfg.LinkCacheDB,
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
