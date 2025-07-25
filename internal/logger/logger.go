package logger

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"stream-radar/internal/config"
	"stream-radar/internal/utils"
	"sync"
)

var logger *zap.Logger
var once sync.Once

var LogLevel = map[string]zapcore.Level{
	"INFO":  zapcore.InfoLevel,
	"ERROR": zapcore.ErrorLevel,
	"DEBUG": zapcore.DebugLevel,
}

func GetInstance() *zap.Logger {
	once.Do(func() {
		logger = initializeLogger(config.LoggerConfig{
			Env:   utils.GetEnv("ENV", "development"),
			Level: utils.GetEnv("LOG_LEVEL", "INFO"),
		})
	})
	return logger
}

func initializeLogger(loggerConfig config.LoggerConfig) *zap.Logger {
	env := loggerConfig.Env

	stdout := zapcore.AddSync(os.Stdout)
	level := zap.NewAtomicLevelAt(LogLevel[loggerConfig.Level])

	var core zapcore.Core
	if env == "production" {
		encoder := ecszap.NewDefaultEncoderConfig()
		core = ecszap.NewCore(encoder, stdout, level)

	} else {
		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder := zapcore.NewConsoleEncoder(developmentCfg)
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, stdout, level),
		)
	}

	return zap.New(core, zap.AddCaller())
}
