package logger

import (
	"log/slog"
	"os"
)

var globalLogger *slog.Logger

func Init(levelDebug slog.Level) {
	globalLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: levelDebug}))
}

func Debug(msg string, fields ...any) {
	globalLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...any) {
	globalLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...any) {
	globalLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...any) {
	globalLogger.Error(msg, fields...)
}

//
//func Fatal(msg string, fields ...any) {
//	globalLogger.Fatal(msg, fields...)
//}
//
//func WithOptions(opts ...zap.Option) *zap.Logger {
//	return globalLogger.WithOptions(opts...)
//}
