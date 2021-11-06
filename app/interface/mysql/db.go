package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"artsign/ent"

	"artsign/ent/migrate"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-sql-driver/mysql"
)

const defaultMySQLPort = 3306

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string

	Loc       string
	Collation string
}

func NewClient(db *sql.DB) (*ent.Client, error) {

	drv := entsql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		return nil, err
	}

	return client, nil
}

func NewDB(cfg Config) (*sql.DB, error) {

	connStr, err := buildConnectionString(cfg)
	if err != nil {
		return nil, err
	}

	return sql.Open("mysql", connStr)
}

func buildConnectionString(cfg Config) (string, error) {
	mysqlCfg := mysql.NewConfig()

	if cfg.Host == "" {
		return "", errors.New("db host is not set")
	}

	if cfg.User == "" {
		return "", errors.New("db user is not set")
	}

	mysqlCfg.Net = "tcp"
	port := defaultMySQLPort
	if cfg.Port != 0 {
		port = cfg.Port
	}

	mysqlCfg.Addr = fmt.Sprintf("%s:%d", cfg.Host, port)

	mysqlCfg.DBName = cfg.Database
	mysqlCfg.User = cfg.User
	mysqlCfg.Passwd = cfg.Password

	mysqlCfg.ParseTime = true

	if cfg.Loc != "" {
		loc, err := time.LoadLocation(cfg.Loc)
		if err != nil {
			return "", err
		}
		mysqlCfg.Loc = loc
	}

	if cfg.Collation != "" {
		mysqlCfg.Collation = cfg.Collation
	}

	ret := mysqlCfg.FormatDSN()

	return ret, nil
}
