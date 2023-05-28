//go:build !tlProduction

package tl

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	if logTopLevel != nil {
		return
	}

	cfg := zap.NewDevelopmentConfig()
	cfg.Level.SetLevel(zapcore.DebugLevel)

	if TUSK_LOG_LEVEL := os.Getenv("LOG_LEVEL"); TUSK_LOG_LEVEL != "" {
		if level, err := zapcore.ParseLevel(TUSK_LOG_LEVEL); err == nil {
			cfg.Level.SetLevel(level)
		}
	}

	if TUSK_LOG_FILE := os.Getenv("LOG_FILE"); TUSK_LOG_FILE != "" {
		cfg.OutputPaths = append(cfg.OutputPaths, TUSK_LOG_FILE)
	}

	log, err := cfg.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	if err != nil {
		panic(err)
	}

	SetLogger(log)
}
