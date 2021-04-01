package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func NewLogger() *logrus.Logger {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	return log
}

// Error exibe detalhes do erro
func Error(args ...interface{}) {
	log.Error(args...)
}

// ErrorContext exibe detalhes do erro com o contexto
func ErrorContext(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Error(args...)
}

// Info exibe detalhes do log info
func Info(args ...interface{}) {
	log.Info(args...)
}

// InfoContext exibe detalhes do log info com o contexto
func InfoContext(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Info(args...)
}

// Debug exibe detalhes do log debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// DebugContext exibe detalhes do log debug com o contexto
func DebugContext(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Debug(args...)
}

// Trace exibe detalhes do log trace
func Trace(args ...interface{}) {
	log.Trace(args...)
}

// TraceContext exibe detalhes do log trace com o contexto
func TraceContext(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Trace(args...)
}

// Fatal exibe detalhes do erro
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
