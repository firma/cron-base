package job

import (
	"context"
	"github.com/robfig/cron/v3"
)

type Job interface {
	RegisterJob(subscribe RegisterJob)
	Start() error
	Stop() context.Context
	Remove(entryId cron.EntryID) error
}

type JobHandler func(ctx context.Context) error

type RegisterJob interface {
	Name() string
	Spec() string
	Handler() JobHandler
}

type jobInfo struct {
	spec string
	name string
	h    JobHandler
}

func (d jobInfo) Name() string {
	return d.name
}

func (d jobInfo) Spec() string {
	return d.spec
}

func (d jobInfo) Handler() JobHandler {
	return d.h
}

func NewJobInfo(name, spec string, handler JobHandler) RegisterJob {
	return jobInfo{
		name: name,
		spec: spec,
		h:    handler,
	}
}
