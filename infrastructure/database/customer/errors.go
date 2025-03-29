package database

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	CustomerNotFound = "customer_not_found"
)

func CustomerNotFoundError(err error) zap.Field {
	return zapcore.Field{
		Key:    "error",
		Type:   zapcore.StringType,
		String: err.Error(),
	}
}
