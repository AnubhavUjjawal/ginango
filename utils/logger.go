package utils

import "go.uber.org/zap"

// TODO: configure logger based on env.
// TODO: Allow for a development logger
var Logger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func init() {
	Logger, _ = zap.NewProduction()
	defer Logger.Sync() // flushes buffer, if any
	SugaredLogger = Logger.Sugar()
}
