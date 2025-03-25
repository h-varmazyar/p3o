package configs

import (
	gormext "github.com/h-varmazyar/gopack/gorm"
)

func New() (Configs, error) {
	cfg, err := load()
	if err != nil {
		return Configs{}, err
	}

	return cfg, nil
}

type Configs struct {
	DB          gormext.Configs
	Redis       Redis
	Server      Server
	UserService UserService
}
