package config

import (
	"os"
	"strconv"
)

type Configs struct {
	Db  DbConfig
	Log LoggerConfig
}

type DbConfig struct {
	AppName  string
	User     string
	Password string
	Host     string
	Port     string
	Db       string
	MaxConn  int32
	MinConn  int32
	IdleTime int
}

type LoggerConfig struct {
	Env   string
	Level string
}

func InitConfigs() *Configs {
	var config Configs

	config.Db = getDbConfig()
	config.Log = getLogConfig()

	return &config

}

func getLogConfig() LoggerConfig {
	return LoggerConfig{
		Env:   os.Getenv("ENV"),
		Level: os.Getenv("LOG_LEVEL"),
	}
}

func getDbConfig() DbConfig {
	maxConn := "1"
	minConn := "1"
	idleTime := "1000"

	maxConnInt, _ := strconv.Atoi(maxConn)
	minConnInt, _ := strconv.Atoi(minConn)
	idleTimeInt, _ := strconv.Atoi(idleTime)

	dbConfig := DbConfig{
		AppName:  os.Getenv("APP_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Db:       os.Getenv("DB_NAME"),
		MaxConn:  int32(maxConnInt),
		MinConn:  int32(minConnInt),
		IdleTime: idleTimeInt,
	}

	return dbConfig
}
