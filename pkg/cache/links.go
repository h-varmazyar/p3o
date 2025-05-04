package cache

import (
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type Link struct {
    URL string
    ID uint32
    OwnerID uint32
    SystemLink bool
}

func (c *RedisCache) Get(key string) (Link, bool, error) {
    val, err := c.client.Get(c.ctx, key).Result()
    if err == redis.Nil {
        return Link{}, false, nil
    } else if err != nil {
        return Link{}, false, err
    }

    link:=Link{}

    err=json.Unmarshal([]byte(val), &link)
    if err != nil {
        return Link{}, false, err
    }

    return link, true, nil
}

func (c *RedisCache) Set(key string, link Link) error {
    cacheValue, err:=json.Marshal(link)
    if err != nil {
        return err
    }
    return c.client.Set(c.ctx, key, cacheValue, 0).Err() // No TTL, LFU handles eviction
}

func (c *RedisCache) Delete(key string) error {
    return c.client.Del(c.ctx, key).Err()
}