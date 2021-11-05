package env

import (
	"artsign/app/interface/mysql"
	"artsign/app/interface/s3"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ArtsignMySQL mysql.Config
	S3           s3.Config
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)

	var m mysql.Config
	err := envconfig.Process("artsign_db", &m)
	if err != nil {
		return nil, err
	}
	cfg.ArtsignMySQL = m

	var s s3.Config
	err = envconfig.Process("s3", &s)
	if err != nil {
		return nil, err
	}
	cfg.S3 = s

	return cfg, nil
}
