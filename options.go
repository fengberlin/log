package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type options struct {
	level zap.AtomicLevel
	encoderCfg zapcore.EncoderConfig
	console bool
	serviceName string
	fileName string
	fileRotateMaxAge time.Duration
	fileRotationTime time.Duration
	development bool
	addCaller bool
	callSkip int
	addStacktrace zapcore.LevelEnabler
	wrapCoreFunc func(zapcore.Core) zapcore.Core
}

type Option interface {
	apply(*options)
}

type optionFunc struct {
	f func(*options)
}

func (fo *optionFunc) apply(opts *options) {
	fo.f(opts)
}

func newOptionFunc(f func(*options)) *optionFunc {
	return &optionFunc{f: f}
}

func WithLevel(l zapcore.Level) Option {
	return newOptionFunc(func(o *options) {
		o.level = zap.NewAtomicLevelAt(l)
	})
}

func WithServiceName(serviceName string) Option {
	return newOptionFunc(func(o *options) {
		o.serviceName = serviceName
	})
}

func WithFileName(fileName string) Option {
	return newOptionFunc(func(o *options) {
		o.fileName = fileName
	})
}