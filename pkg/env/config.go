package env

import (
	"artsign/app/interface/mysql"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ArtsignMySQL mysql.Config
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)

	var m mysql.Config
	err := envconfig.Process("artsign_db", &m)
	if err != nil {
		return nil, err
	}
	cfg.ArtsignMySQL = m

	return cfg, nil
}
