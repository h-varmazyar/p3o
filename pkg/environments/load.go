package environments

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadEnvironments(configs interface{}) error {
	if err := godotenv.Load("./configs/.env"); err != nil {
		return err
	}
	if err := env.Parse(configs); err != nil {
		return err
	}
	return nil
}
