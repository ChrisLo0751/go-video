package initialize

import "go.uber.org/zap"

func InitialLogger() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
