package lib

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var instance *zap.Logger

func InitLogger() {
	var err error

	logPath, err := getLogPath()
	if err != nil {
		log.Fatal(err)
	}

	fileWriter := zapcore.AddSync(&lumberjackWriter{path: logPath})

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		NameKey:      "logger",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		fileWriter,
		zapcore.DebugLevel,
	)

	instance = zap.New(
		core,
		zap.AddCaller(),
	)
}

func getLogPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	logDir := filepath.Join(configDir, "plainslate", "debug")
	err = os.MkdirAll(logDir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(logDir, "debug.log"), nil
}

type lumberjackWriter struct {
	path string
}

func (l *lumberjackWriter) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(l.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(p)
}

func CloseLogger() {
	if instance != nil {
		_ = instance.Sync() // Flush logs
	}
}

func CommonLog(ctx context.Context, level, message string, fields ...zap.Field) {
	switch level {
	case "debug":
		instance.Debug(message, fields...)
	case "info":
		instance.Info(message, fields...)
	case "warn":
		instance.Warn(message, fields...)
	case "error":
		instance.Error(message, fields...)
	case "panic":
		instance.Panic(message, fields...)
	case "fatal":
		instance.Fatal(message, fields...)
	}
}

func LogDebug(ctx context.Context, message string, fields ...zap.Field) {
	fields = append(fields, zap.String("component", "backend"))
	CommonLog(ctx, "debug", message, fields...)
}

func LogInfo(ctx context.Context, message string, fields ...zap.Field) {
	fields = append(fields, zap.String("component", "backend"))
	CommonLog(ctx, "info", message, fields...)
}

func LogWarn(ctx context.Context, message string, fields ...zap.Field) {
	fields = append(fields, zap.String("component", "backend"))
	CommonLog(ctx, "warn", message, fields...)
}

func LogError(ctx context.Context, message string, fields ...zap.Field) {
	fields = append(fields, zap.String("component", "backend"))
	CommonLog(ctx, "error", message, fields...)
}

func LogPanic(ctx context.Context, message string, fields ...zap.Field) {
	fields = append(fields, zap.String("component", "backend"))
	CommonLog(ctx, "panic", message, fields...)
}

func LogFatal(ctx context.Context, message string, fields ...zap.Field) {
	fields = append(fields, zap.String("component", "backend"))
	CommonLog(ctx, "fatal", message, fields...)
}
