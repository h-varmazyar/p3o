package cache

import (
	"context"
	"encoding/json"
	"github.com/h-varmazyar/p3o/configs"
	"github.com/h-varmazyar/p3o/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

type Link struct {
	URL     string
	ID      uint
	OwnerID uint
	Error   *errors.Error
}

type LinkRedisCache struct {
	log    *log.Logger
	client *redis.Client
	ctx    context.Context
}

func NewLinkRedisCache(log *log.Logger, cfg configs.Redis) (*LinkRedisCache, error) {
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

	if err := client.ConfigSet(ctx, "maxmemory-policy", "allkeys-lfu").Err(); err != nil {
		return nil, err
	}

	return &LinkRedisCache{
		log:    log,
		client: client,
		ctx:    ctx,
	}, nil
}

func (c *LinkRedisCache) Get(key string) (Link, bool, error) {

	val, err := c.client.Get(c.ctx, key).Result()
	if err == redis.Nil {
		return Link{}, false, nil
	} else if err != nil {
		return Link{}, false, err
	}

	link := Link{}

	err = json.Unmarshal([]byte(val), &link)
	if err != nil {
		return Link{}, false, err
	}

	return link, true, nil
}

func (c *LinkRedisCache) Set(key string, link Link) error {
	cacheValue, err := json.Marshal(link)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, key, cacheValue, 0).Err() // No TTL, LFU handles eviction
}

func (c *LinkRedisCache) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}
