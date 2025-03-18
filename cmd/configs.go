package main

import (
	gormext "github.com/h-varmazyar/gopack/gorm"
	"github.com/h-varmazyar/p3o/internal/repositories/link"
)

type Configs struct {
	Version       string
	HttpAddress   string
	RedisAddress  string
	RedisUsername string
	RedisPassword string
	LinkCacheDB   int
	DB            *gormext.Configs
	LinkApp       *link.Configs
}

type PublicConfigs struct {
	RedisConfigs *RedisConfigs
}

type RedisConfigs struct {
}
