package main

import (
	"go.uber.org/zap"
)

const version = "0.1.0"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("start watchdog",
		zap.String("version", version),
	)
}
