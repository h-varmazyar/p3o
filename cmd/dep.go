package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/configs"
	authController "github.com/h-varmazyar/p3o/internal/controllers/auth"
	dashboardController "github.com/h-varmazyar/p3o/internal/controllers/dashboard"
	linkController "github.com/h-varmazyar/p3o/internal/controllers/link"
	userController "github.com/h-varmazyar/p3o/internal/controllers/user"
	linkRepository "github.com/h-varmazyar/p3o/internal/repositories/link"
	userRepository "github.com/h-varmazyar/p3o/internal/repositories/user"
	visitRepository "github.com/h-varmazyar/p3o/internal/repositories/visit"
	"github.com/h-varmazyar/p3o/internal/router"
	"github.com/h-varmazyar/p3o/internal/router/middlewares"
	v1 "github.com/h-varmazyar/p3o/internal/router/v1"
	linkService "github.com/h-varmazyar/p3o/internal/services/link"
	userService "github.com/h-varmazyar/p3o/internal/services/user"
	"github.com/h-varmazyar/p3o/internal/workers"
	"github.com/h-varmazyar/p3o/pkg/cache"
	"github.com/h-varmazyar/p3o/pkg/logger"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

type dependencies struct {
	ctx          context.Context
	StopSignal   chan os.Signal
	Log          *log.Logger
	Cfg          configs.Configs
	DB           *gorm.DB
	Redis        *redis.Client
	Gin          *gin.Engine
	VisitChan    chan workers.VisitRecord
	Repositories struct {
		User  userRepository.Repository
		Link  linkRepository.Repository
		Visit visitRepository.Repository
	}
	Services struct {
		User userService.Service
		Link linkService.Service
	}
	Controllers struct {
		AuthController      authController.Controller
		LinkController      linkController.Controller
		DashboardController dashboardController.Controller
		UserController      userController.Controller
	}
	Routers struct {
		Router      router.Router
		V1          v1.Router
		Middlewares struct {
			PublicAuth middlewares.PublicAuthMiddleware
		}
	}
	Cache struct {
		LinkCache             *cache.LinkRedisCache
		VerificationCodeCache *cache.VerificationCodeRedisCache
	}
}

var generalDependencies = func(dep *dependencies) (err error) {
	dep.ctx = context.Background()
	dep.StopSignal = make(chan os.Signal, 1)
	dep.Log = logger.NewLogger()
	dep.Cfg, err = configs.New()
	if err != nil {
		return
	}

	return
}

var infraDependencies = func(dep *dependencies) (err error) {
	dep.DB, err = initializePostgres(dep.Cfg.DB)
	if err != nil {
		return
	}
	dep.Redis = initializeRedis(dep.Cfg.Redis)
	dep.Gin = initializeGin(dep.Log)
	dep.VisitChan = initializeVisitChannel()

	return
}

var controllerDependencies = func(dep *dependencies) (err error) {
	dep.Controllers.AuthController = authController.New(dep.Services.User)
	dep.Controllers.LinkController = linkController.New(dep.Services.Link, dep.VisitChan)
	dep.Controllers.DashboardController = dashboardController.New(dep.Services.Link)
	return
}

var repositoryDependencies = func(dep *dependencies) (err error) {
	dep.Repositories.User = userRepository.New(dep.Log, dep.DB)
	dep.Repositories.Link = linkRepository.New(dep.Log, dep.DB)
	dep.Repositories.Visit = visitRepository.New(dep.Log, dep.DB)
	return
}

var cacheDependencies = func(dep *dependencies) (err error) {
	dep.Cache.LinkCache, err = cache.NewLinkRedisCache(dep.Log, dep.Cfg.Redis)
	if err != nil {
		return
	}
	dep.Cache.VerificationCodeCache, err = cache.NewVerificationCodeRedisCache(dep.Log, dep.Cfg.Redis)
	if err != nil {
		return
	}
	return
}

var serviceDependencies = func(dep *dependencies) (err error) {
	dep.Services.User, err = userService.New(dep.Log, dep.Cfg.UserService, dep.Repositories.User, dep.Cache.VerificationCodeCache)
	if err != nil {
		return
	}

	dep.Services.Link = linkService.New(dep.Log, dep.Cfg.LinkService, dep.Repositories.Link, dep.Repositories.Visit, dep.Cache.LinkCache)
	return
}

var routersDependencies = func(dep *dependencies) (err error) {
	dep.Routers.Middlewares.PublicAuth = middlewares.NewPublicAuthMiddleware(dep.Log)
	dep.Routers.V1 = v1.New(dep.Controllers.AuthController, dep.Controllers.LinkController, dep.Controllers.DashboardController, dep.Controllers.UserController, dep.Routers.Middlewares.PublicAuth)
	dep.Routers.Router = router.New(dep.Log, dep.Routers.V1, dep.Services.Link)
	return
}

func InjectDependencies() (dep dependencies, err error) {
	if err = generalDependencies(&dep); err != nil {
		return
	}

	if err = infraDependencies(&dep); err != nil {
		return
	}

	if err = cacheDependencies(&dep); err != nil {
		return
	}

	if err = repositoryDependencies(&dep); err != nil {
		return
	}

	if err = serviceDependencies(&dep); err != nil {
		return
	}

	if err = controllerDependencies(&dep); err != nil {
		return
	}

	if err = routersDependencies(&dep); err != nil {
		return
	}

	dep.Log.Info("Injecting dependencies")

	return
}
