package glog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

type Field = zapcore.Field

var Int64 = zap.Int64
var Int32 = zap.Int32
var Int = zap.Int
var Uint32 = zap.Uint32
var String = zap.String
var Any = zap.Any
var Err = zap.Error
var NamedErr = zap.NamedError
var Bool = zap.Bool
var Duration = zap.Duration

type Config struct {
	EnableConsole bool
	ConsoleJson   bool
	ConsoleLevel  string
}

type Logger struct {
	zap          *zap.Logger
	consoleLevel zap.AtomicLevel
}

func Init(config Config) *Logger {
	cores := []zapcore.Core{}
	logger := &Logger{
		consoleLevel: zap.NewAtomicLevelAt(getZapLevel(config.ConsoleLevel)),
	}

	if config.EnableConsole {
		writer := zapcore.Lock(os.Stderr)
		core := zapcore.NewCore(makeEncoder(config.ConsoleJson), writer, logger.consoleLevel)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger.zap = zap.New(
		combinedCore,
		zap.AddCaller(),
	)

	return logger
}

func makeEncoder(json bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	if json {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case DebugLevel:
		return zapcore.DebugLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func (l *Logger) SetConsoleLevel(level string) {
	l.consoleLevel.SetLevel(getZapLevel(level))
}

func (l *Logger) With(fields ...Field) *Logger {
	newlogger := *l
	newlogger.zap = newlogger.zap.With(fields...)
	return &newlogger
}

func (l *Logger) Debug(message string, fields ...Field) {
	l.zap.Debug(message, fields...)
}

func (l *Logger) Info(message string, fields ...Field) {
	l.zap.Info(message, fields...)
}

func (l *Logger) Warn(message string, fields ...Field) {
	l.zap.Warn(message, fields...)
}

func (l *Logger) Error(message string, fields ...Field) {
	l.zap.Error(message, fields...)
}

func (l *Logger) Critical(message string, fields ...Field) {
	l.zap.Error(message, fields...)
}
