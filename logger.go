package log

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {

	l := zap.New(nil)
	return l, nil
}