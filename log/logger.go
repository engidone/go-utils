package log

import (
	"log"
	"os"
	"time"

	"github.com/engidone/utils/clog"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger
var isColored bool

func init() {
	isColored = os.Getenv("LOGGING") == "console"
	if isColored {
		return
	}

	baseLogger, err := zap.NewProduction(zap.AddCaller())
	if err != nil {
		log.Fatal(err)
	}
	logger := baseLogger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	sugar = logger
	defer logger.Sync()
}

func Info(args ...any) {
	if isColored {
		clog.Info(args...)
		return
	}
	sugar.Info(args...)
}

func Success(args ...any) {
	if isColored {
		clog.Success(args...)
		return
	}
	sugar.Info(args...)
}

func Error(args ...any) {
	if isColored {
		clog.Error(args...)
		return
	}
	sugar.Error(args...)
}

func Warn(args ...any) {
	if isColored {
		clog.Warn(args...)
		return
	}
	sugar.Warn(args...)
}

func Fatal(args ...any) {
	if isColored {
		clog.Fatal(args...)
		return
	}
	sugar.Fatal(args...)
}

func Debug(args ...any) {
	if isColored {
		clog.Debug(args...)
		return
	}
	sugar.Debug(args...)
}

func Infof(template string, args ...any) {
	if isColored {
		clog.Infof(template, args...)
		return
	}
	sugar.Infof(template, args...)
}

func Successf(template string, args ...any) {
	if isColored {
		clog.Successf(template, args...)
		return
	}
	sugar.Infof(template, args...)
}

func Fatalf(template string, args ...any) {
	if isColored {
		clog.Fatalf(template, args...)
		return
	}
	sugar.Fatalf(template, args...)
}

func Errorf(template string, args ...any) {
	if isColored {
		clog.Errorf(template, args...)
		return
	}
	sugar.Errorf(template, args...)
}

func Warnf(template string, args ...any) {
	if isColored {
		clog.Warnf(template, args...)
		return
	}
	sugar.Warnf(template, args...)
}

func Debugf(template string, args ...any) {
	if isColored {
		clog.Debugf(template, args...)
		return
	}
	sugar.Debugf(template, args...)
}

func Infow(msg string, keysAndValues ...any) {
	sugar.Infow(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...any) {
	sugar.Errorw(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...any) {
	sugar.Warnw(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...any) {
	sugar.Debugw(msg, keysAndValues...)
}

type LogField = zap.Field

func String(key, value string) LogField {
	return zap.String(key, value)
}

func Int(key string, value int) LogField {
	return zap.Int(key, value)
}

func Int32(key string, value int32) LogField {
	return zap.Int32(key, value)
}

func Bool(key string, value bool) LogField {
	return zap.Bool(key, value)
}

func Err(err error) LogField {
	return zap.Error(err)
}

func Any(key string, value any) LogField {
	return zap.Any(key, value)
}

func Float64(key string, value float64) LogField {
	return zap.Float64(key, value)
}

func Int64(key string, value int64) LogField {
	return zap.Int64(key, value)
}

func Uint(key string, value uint) LogField {
	return zap.Uint(key, value)
}

func Uint32(key string, value uint32) LogField {
	return zap.Uint32(key, value)
}

func Uint64(key string, value uint64) LogField {
	return zap.Uint64(key, value)
}

func Duration(key string, value time.Duration) LogField {
	return zap.Duration(key, value)
}

func Time(key string, value time.Time) LogField {
	return zap.Time(key, value)
}
