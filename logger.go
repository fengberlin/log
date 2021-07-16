package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type innerLogger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger
	once          sync.Once
}

const (
	defaultTimeFormatLayout = "2006-01-02T15:04:05.000Z07:00"
)

// defaultTimeEncoder init a zap customized time encoder
func defaultTimeEncoder() zapcore.TimeEncoder {
	//enc.AppendString(t.Format(defaultTimeFormatLayout))
	return zapcore.TimeEncoderOfLayout(defaultTimeFormatLayout)
}

func defaultEncoderConfig() zapcore.EncoderConfig {
	cfg := zapcore.EncoderConfig{
		TimeKey:          "ts",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      zapcore.OmitKey,
		MessageKey:       "msg",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       defaultTimeEncoder(),
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: "",
	}
	return cfg
}

func NewLogger(opts ...Option) (*zap.Logger, error) {
	l := zap.New(nil)
	return l, nil
}
