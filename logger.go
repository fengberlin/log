package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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
		ConsoleSeparator: "", // default is a tab
	}
	return cfg
}

func defaultLogOptions() *options {
	return &options{
		level:            zap.AtomicLevel{},
		encoderCfg:       zapcore.EncoderConfig{},
		serviceName:      "",
		fileName:         "",
		fileRotateMaxAge: 0,
		fileRotationTime: 0,
		addCaller:        false,
		callSkip:         0,
		addStacktrace:    nil,
		wrapCoreFunc:     nil,
	}
}

func NewLogger(opts ...Option) (*zap.Logger, error) {
	defaultOpts := defaultLogOptions()
	defaultOpts.encoderCfg = defaultEncoderConfig()
	for i := 0; i < len(opts); i++ {
		opts[i].apply(defaultOpts)
	}
	var encoder zapcore.Encoder
	var cores []zapcore.Core
	if defaultOpts.serviceName == "" && defaultOpts.fileName == "" {
		encoder = zapcore.NewConsoleEncoder(defaultOpts.encoderCfg)
		cores = []zapcore.Core{zapcore.NewCore(encoder, os.Stderr, defaultOpts.level)}
	}
	l := zap.New(zapcore.NewTee(cores...))
	return l, nil
}
