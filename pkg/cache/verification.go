package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

const expirationDuration = time.Minute * 2

type VerificationCodeRedisCache struct {
	log    *log.Logger
	client *redis.Client
	ctx    context.Context
}

func NewVerificationCodeRedisCache(log *log.Logger, cfg configs.Redis) (*VerificationCodeRedisCache, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.LinkCacheDB,
		PoolSize: 10,
	})

	// Test connection
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	if err := client.ConfigSet(ctx, "maxmemory", "100mb").Err(); err != nil {
		return nil, err
	}

	return &VerificationCodeRedisCache{
		log:    log,
		client: client,
		ctx:    ctx,
	}, nil
}

func (c *VerificationCodeRedisCache) Get(userId uint) (string, error) {
	val, err := c.client.Get(c.ctx, fmt.Sprint(userId)).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (c *VerificationCodeRedisCache) Set(userId uint, code string) error {
	return c.client.SetEx(c.ctx, fmt.Sprint(userId), code, expirationDuration).Err() // No TTL, LFU handles eviction
}
