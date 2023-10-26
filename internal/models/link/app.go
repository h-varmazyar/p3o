package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/models/link/repository"
	"github.com/h-varmazyar/p3o/internal/models/link/service"
	"github.com/h-varmazyar/p3o/internal/models/link/workers"
	db "github.com/h-varmazyar/p3o/pkg/db/PostgreSQL"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type App struct {
	service service.Service
}

func NewApp(ctx context.Context, log *log.Logger, db *db.DB, redisClient *redis.Client, configs *Configs) (*App, error) {
	visitChannel := make(chan workers.VisitRecord, 1000)
	persistDB, err := repository.NewPostgresRepository(ctx, log, db)
	if err != nil {
		return nil, err
	}

	cacheDB, err := repository.NewRedisRepository(ctx, log, redisClient, configs.LinkCacheTTL)
	if err != nil {
		return nil, err
	}

	err = workers.StartVisitWorker(log, visitChannel, persistDB)
	if err != nil {
		return nil, err
	}

	s := service.NewLinkService(log, persistDB, cacheDB, visitChannel)

	return &App{service: s}, nil
}

func (a *App) GetService() service.Service {
	return a.service
}
