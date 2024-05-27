package redis

type Configs struct {
	RedisAddress  string `yaml:"redisAddress"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPassword"`
	LinkCacheDB   int    `yaml:"linkCacheDB"`
}
