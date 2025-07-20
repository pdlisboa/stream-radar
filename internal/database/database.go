package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/url"
	"stream-radar/internal/config"
	"stream-radar/internal/logger"
	"time"
)

type DbPool interface {
	QueryRow(context.Context, string, ...any) pgx.Row
	Query(context.Context, string, ...any) (pgx.Rows, error)
	Close()
	Ping(context.Context) error
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

var Pool DbPool

func Connect(dbConfig config.DbConfig) (DbPool, error) {
	log := logger.GetInstance()
	pgURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(dbConfig.User, dbConfig.Password),
		Host:   dbConfig.Host + ":" + dbConfig.Port,
		Path:   "/" + dbConfig.Db,
	}

	dsn := pgURL.String()
	config, dbConfigError := pgxpool.ParseConfig(dsn)

	if dbConfigError != nil {
		log.Fatal("Error parsing database config")
		return nil, dbConfigError
	}

	config.MaxConns = dbConfig.MaxConn
	config.MinConns = dbConfig.MinConn
	config.MaxConnIdleTime = time.Minute * time.Duration(dbConfig.IdleTime)
	config.ConnConfig.RuntimeParams["application_name"] = dbConfig.AppName

	Pool, _ = pgxpool.NewWithConfig(context.Background(), config)

	pingErr := Pool.Ping(context.Background())
	if pingErr != nil {
		log.Fatal("unable to connect to database")
		panic("unable to connect to database")
	}

	log.Info("Database connected!")
	return Pool, nil
}
