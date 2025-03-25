package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func loadString(envName string) string {
	validate(envName)
	return viper.GetString(envName)
}

func loadInt(envName string) int {
	validate(envName)
	return viper.GetInt(envName)
}

func loadInt64(envName string) int64 {
	validate(envName)
	return viper.GetInt64(envName)
}

func loadUint(envName string) uint {
	validate(envName)
	return viper.GetUint(envName)
}

func loadUint64(envName string) uint64 {
	validate(envName)
	return viper.GetUint64(envName)
}

func loadBool(envName string) bool {
	validate(envName)
	return viper.GetBool(envName)
}

//func loadUint(envName string) uint {
//	validate(envName)
//	return viper.GetUint(envName)
//}

func loadFloat(envName string) float64 {
	validate(envName)
	return viper.GetFloat64(envName)
}

func loadStringSlice(envName string) []string {
	validate(envName)
	return viper.GetStringSlice(envName)
}

// nolint:unused
func loadDuration(envName string) time.Duration {
	validate(envName)
	return viper.GetDuration(envName)
}

func validate(envName string) {
	exists := viper.IsSet(envName)
	if !exists {
		panic(fmt.Sprintf("environment variable [%s] does not exist", envName))
	}
}
