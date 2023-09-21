package logs

import (
	"context"
	"go.uber.org/zap"
)

func NameSpace(name string) *zap.Logger {
	return logger.Named(name)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func CtxDebug(ctx context.Context, msg string, fields ...zap.Field) {
	ctxFields := GetAllFields(ctx)
	logger.With(ctxFields...).Debug(msg, fields...)
}

func CtxInfo(ctx context.Context, msg string, fields ...zap.Field) {
	ctxFields := GetAllFields(ctx)
	logger.With(ctxFields...).Info(msg, fields...)
}

func CtxWarn(ctx context.Context, msg string, fields ...zap.Field) {
	ctxFields := GetAllFields(ctx)
	logger.With(ctxFields...).Warn(msg, fields...)
}

func CtxError(ctx context.Context, msg string, fields ...zap.Field) {
	ctxFields := GetAllFields(ctx)
	logger.With(ctxFields...).Error(msg, fields...)
}
