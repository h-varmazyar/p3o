package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/configs"
	authController "github.com/h-varmazyar/p3o/internal/controllers/auth"
	linkController "github.com/h-varmazyar/p3o/internal/controllers/link"
	linkRepository "github.com/h-varmazyar/p3o/internal/repositories/link"
	userRepository "github.com/h-varmazyar/p3o/internal/repositories/user"
	"github.com/h-varmazyar/p3o/internal/router"
	v1 "github.com/h-varmazyar/p3o/internal/router/v1"
	linkService "github.com/h-varmazyar/p3o/internal/services/link"
	userService "github.com/h-varmazyar/p3o/internal/services/user"
	"github.com/h-varmazyar/p3o/internal/workers"
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
		User userRepository.Repository
		Link linkRepository.Repository
	}
	Services struct {
		User userService.Service
		Link linkService.Service
	}
	Controllers struct {
		AuthController authController.Controller
		LinkController linkController.Controller
	}
	Routers struct {
		Router router.Router
		V1     v1.Router
	}
}

var generalDependencies = func(dep *dependencies) (err error) {
	dep.ctx = context.Background()
	dep.StopSignal = make(chan os.Signal, 1)
	dep.Log = log.New()
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
	return
}

var repositoryDependencies = func(dep *dependencies) (err error) {
	dep.Repositories.User = userRepository.New(dep.Log, dep.DB)
	dep.Repositories.Link = linkRepository.New(dep.Log, dep.DB)

	return
}

var serviceDependencies = func(dep *dependencies) (err error) {
	dep.Services.User, err = userService.New(dep.Log, dep.Cfg.UserService, dep.Repositories.User)
	if err != nil {
		return
	}

	dep.Services.Link = linkService.New(dep.Log, dep.Repositories.Link)
	return
}

var routersDependencies = func(dep *dependencies) (err error) {
	dep.Routers.V1 = v1.New(dep.Controllers.AuthController, dep.Controllers.LinkController)
	dep.Routers.Router = router.New(dep.Log, dep.Routers.V1, dep.Services.Link, dep.VisitChan)
	return
}

func InjectDependencies() (dep dependencies, err error) {
	if err = generalDependencies(&dep); err != nil {
		return
	}

	if err = infraDependencies(&dep); err != nil {
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

	return
}
