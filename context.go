package context

import (
	"context"
	"time"
)

type Context interface {
	context.Context
	Log() Logger
}

type Logger interface {
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
}
type loggerContext struct {
	log Logger
}

func (c *loggerContext) Log() Logger {
	return c.log
}
func (*loggerContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*loggerContext) Done() <-chan struct{} {
	return nil
}

func (*loggerContext) Err() error {
	return nil
}

func (*loggerContext) Value(key interface{}) interface{} {
	return nil
}

func WithLogger(logger Logger) Context {
	return &loggerContext{log: logger}
}
