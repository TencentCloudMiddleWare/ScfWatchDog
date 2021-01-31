package main

import (
	"os"
	"strconv"

	"github.com/TencentCloudMiddleWare/ScfWatchDog/pkg/watchdog"
	"go.uber.org/zap"
)

func main() {
	isdebug, _ := strconv.ParseBool(os.Getenv("WATCHDOG_DEBUG"))
	loglevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	if isdebug {
		loglevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	logger, _ := zap.Config{
		Level:       loglevel,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	watch := watchdog.New(logger)
	watch.Run()

}
