package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

const expirationDuration = time.Minute * 2

type VerificationCode struct {
	Code   string `json:"code"`
	Mobile string `json:"mobile"`
}
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

func (c *VerificationCodeRedisCache) Get(userId uint) (VerificationCode, error) {
	val, err := c.client.Get(c.ctx, fmt.Sprint(userId)).Result()
	if errors.Is(err, redis.Nil) {
		return VerificationCode{}, nil
	} else if err != nil {
		return VerificationCode{}, err
	}

	vc := VerificationCode{}
	if err = json.Unmarshal([]byte(val), &vc); err != nil {
		return VerificationCode{}, err
	}

	return vc, nil
}

func (c *VerificationCodeRedisCache) Set(userId uint, codeValue VerificationCode) error {
	encoded, err := json.Marshal(codeValue)
	if err != nil {
		return err
	}
	return c.client.SetEx(c.ctx, fmt.Sprint(userId), string(encoded), expirationDuration).Err() // No TTL, LFU handles eviction
}
