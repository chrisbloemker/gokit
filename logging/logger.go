package logging

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

// Convenience wrapper for structured logging
func ZapString(key, value string) zap.Field {
	return zap.String(key, value)
}

func ZapError(err error) zap.Field {
	return zap.Error(err)
}
