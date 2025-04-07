package configs

import (
	"errors"
	"fmt"
	gormext "github.com/h-varmazyar/gopack/gorm"
	"github.com/spf13/viper"
)

func load() (Configs, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AllowEmptyEnv(true)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return Configs{}, fmt.Errorf("reading config: %w", err)
		}
	}

	return Configs{
		DB: gormext.Configs{
			Port:        uint16(loadInt("DB_PORT")),
			Host:        loadString("DB_HOST"),
			Username:    loadString("DB_USERNAME"),
			Password:    loadString("DB_PASSWORD"),
			Name:        loadString("DB_NAME"),
			IsSSLEnable: loadBool("DB_SSL_ENABLE"),
		},
		Redis: Redis{
			Address:     loadString("REDIS_ADDRESS"),
			Username:    loadString("REDIS_USERNAME"),
			Password:    loadString("REDIS_PASSWORD"),
			LinkCacheDB: loadInt("REDIS_LINK_CACHE_DB"),
		},
		Server: Server{
			HttpAddress: loadString("SERVER_HOST"),
			HttpPort:    loadInt("SERVER_PORT"),
		},
		UserService: UserService{
			//JWTPublicKey:  loadString("JWT_PUBLIC_KEY"),
			//JWTPrivateKey: loadString("JWT_PRIVATE_KEY"),
		},
		LinkService: LinkService{
			IndirectBaseURL: loadString("INDIRECT_BASE_URL"),
		},
	}, nil
}
