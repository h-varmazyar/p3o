package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/h-varmazyar/p3o/configs"
)

type RedisCache struct {
    client *redis.Client
    ctx    context.Context
}

func NewRedisCache(addr, cfg configs.Redis) (*RedisCache, error) {
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

	if err := client.ConfigSet(ctx, "maxmemory", "100mb").Err();err != nil {
		return nil, err
	}

	if err := client.ConfigSet(ctx, "maxmemory-policy", "allkeys-lfu").Err(); err != nil {
		return nil, err
	}

    return &RedisCache{
        client: client,
        ctx:    ctx,
    }, nil
}