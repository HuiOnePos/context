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
	Infow(msg string, args ...interface{})
	Warnw(msg string, args ...interface{})
	Errorw(msg string, args ...interface{})
	Debugw(msg string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}
type childContext struct {
	context.Context
}

func (c *childContext) Log() Logger {
	return nil
}

func FromSysContext(ctx context.Context) Context {
	return &childContext{Context: ctx}
}

type loggerContext struct {
	Context
	log Logger
}

func (c *loggerContext) Log() Logger {
	return c.log
}
func (c *loggerContext) Deadline() (deadline time.Time, ok bool) {
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

func WithLogger(logger Logger, parent Context) Context {
	return &loggerContext{Context: parent, log: logger}
}
