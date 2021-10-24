package log

import (
	"context"
	"fmt"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log struct {
	logProvider *zap.Logger
}

const (
	skip = 2
)

var logInstance *log

func NewLogger() error {
	if logInstance == nil {
		var opt []zap.Option
		logConfig := zap.NewProductionConfig()
		logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		logConfig.DisableCaller = true
		provider, err := logConfig.Build(opt...)
		if err != nil {
			return err
		}
		logInstance = &log{
			logProvider: provider,
		}
	}
	return nil
}

func Error(ctx context.Context, msg string, err error, fields ...Field) {
	file, line := getFileAndLine()
	fileNLine := fmt.Sprintf("%s:%d", file, line)
	fileField := Field{
		zap.String("file", fileNLine),
	}
	errField := Field{
		zap.String("err", err.Error()),
	}
	fields = append(fields, fileField, errField)
	logInstance.logProvider.Error(msg, convertFields(fields)...)
}

func Info(ctx context.Context, msg string, fields ...Field) {
	file, line := getFileAndLine()
	fileNLine := fmt.Sprintf("%s:%d", file, line)
	fileField := Field{
		zap.String("file", fileNLine),
	}
	fields = append(fields, fileField)
	logInstance.logProvider.Info(msg, convertFields(fields)...)
}

func Warn(ctx context.Context, msg string, fields ...Field) {
	file, line := getFileAndLine()
	fileNLine := fmt.Sprintf("%s:%d", file, line)
	fileField := Field{
		zap.String("file", fileNLine),
	}
	fields = append(fields, fileField)
	logInstance.logProvider.Warn(msg, convertFields(fields)...)
}

func getFileAndLine() (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}

	return file, line
}
