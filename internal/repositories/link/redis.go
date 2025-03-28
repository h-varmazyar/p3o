package link

import (
	"context"
	"encoding/json"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

type redisRepository struct {
	client  *redis.Client
	log     *log.Logger
	linkTTL time.Duration
}

//func NewRedisRepository(_ context.Context, logger *log.Logger, redisClient *redis.Client, linkTTL time.Duration) (Model, error) {
//	return &redisRepository{
//		client:  redisClient,
//		log:     logger,
//		linkTTL: linkTTL,
//	}, nil
//}

func (r *redisRepository) Create(ctx context.Context, link *entities.Link) error {
	bin, err := json.Marshal(link)
	if err != nil {
		return errors.ErrCacheInsertFailed.AddOriginalError(err)
	}
	err = r.client.Set(ctx, link.Key, bin, r.linkTTL).Err()
	if err != nil {
		return errors.ErrCacheInsertFailed.AddOriginalError(err)
	}
	return nil
}

func (r *redisRepository) TotalCounts(_ context.Context) (int64, error) {
	return 0, errors.ErrUnimplemented
}

func (r *redisRepository) TotalVisits(_ context.Context) (int64, error) {
	return 0, errors.ErrUnimplemented
}

func (r *redisRepository) ReturnByKey(ctx context.Context, key string) (*entities.Link, error) {
	exp := r.client.Get(ctx, key)
	if exp.Err() != nil {
		return nil, errors.ErrCacheFetchFailed.AddOriginalError(exp.Err())
	}
	link := new(entities.Link)

	result, err := exp.Result()
	if err != nil {
		return nil, errors.ErrCacheFetchFailed.AddOriginalError(err)
	}
	err = json.Unmarshal([]byte(result), link)
	if err != nil {
		return nil, errors.ErrCacheFetchFailed.AddOriginalError(err)
	}

	return link, nil
}

func (r *redisRepository) Visit(ctx context.Context, id uint) error {
	return errors.ErrUnimplemented
}
