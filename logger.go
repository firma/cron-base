package job

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

var loggerTag = "cronjob"

type logger struct {
	ctx context.Context
}

func newLogger(ctx context.Context) cron.Logger {
	return logger{ctx: ctx}
}

func (l logger) Info(msg string, keysAndValues ...any) {
	log.Infof(loggerTag+msg, keysAndValues...)
}
func (l logger) Infof(msg string, keysAndValues ...any) {
	log.Infof(loggerTag+msg, keysAndValues...)
}

func (l logger) Error(err error, msg string, keysAndValues ...any) {
	log.Errorf(loggerTag+msg+err.Error(), keysAndValues...)
}

func (l logger) Errorf(err error, msg string, keysAndValues ...any) {
	log.Errorf(loggerTag+msg+err.Error(), keysAndValues...)
}
