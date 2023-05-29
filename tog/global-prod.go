//go:build tlProduction

package tl

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	if logTopLevel != nil {
		return
	}

	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zapcore.InfoLevel)

	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	SetLogger(log)
}
