package vender

import (
	"database/sql"
	"fmt"
	"phalanx"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type (
	Database  = *bun.DB
	BaseModel = bun.BaseModel
)

type DatabaseConfig interface {
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DB       string
}

type SQLiteConfig struct {
	Filename string
}

func (config PostgresConfig) DSN() string {
	const postgres = "postgres"
	user := func() string {
		if config.User == "" {
			return postgres
		}
		return config.User
	}
	host := func() string {
		if config.User == "" {
			return postgres
		}
		return config.Host
	}
	port := func() int {
		if port := config.Port; port <= 0 || port > 65535 {
			return 5432
		}
		return config.Port
	}
	db := func() string {
		if config.DB == "" {
			return postgres
		}
		return config.DB
	}
	const format = "%s://%s:%s@%s:%d/%s?sslmode=disable"
	return fmt.Sprintf(format, postgres, user(), config.Password, host(), port(), db())
}

func OpenDatabase(config DatabaseConfig) (Database, error) {
	switch config := config.(type) {
	case nil:
		return OpenDatabase(PostgresConfig{})
	case PostgresConfig:
		dsn := config.DSN()
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		db := bun.NewDB(sqldb, pgdialect.New())
		return db, nil
	case SQLiteConfig:
		if config.Filename == "" {
			return nil, phalanx.ErrorRequired("config.Filename")
		}
		sqldb, err := sql.Open(sqliteshim.ShimName, config.Filename)
		if err != nil {
			return nil, err
		}
		return bun.NewDB(sqldb, sqlitedialect.New()), nil
	}
	return nil, phalanx.ErrorUnsupport("config")
}
