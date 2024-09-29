package job

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"time"
)

var _ Job = (*StartJob)(nil)

type StartJob struct {
	cron       *cron.Cron
	jobs       map[string]RegisterJob
	jobEntryId []cron.EntryID
}

func NewJob() Job {
	return StartJob{
		cron: cron.New(cron.WithParser(cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor))),
		jobs: make(map[string]RegisterJob),
	}
}
func (n StartJob) RegisterJob(job RegisterJob) {
	n.jobs[job.Spec()] = job
}

func (n StartJob) Remove(entryId cron.EntryID) error {
	n.cron.Remove(entryId)
	log.Infof("关闭定时任务")
	return nil
}

func (n StartJob) Start() error {

	for _, v := range n.jobs {
		entryId, err := n.cron.AddJob(v.Spec(), warpJob(fmt.Sprintf("%s: %s", v.Name(), v.Spec()), v.Handler()))
		if err != nil {
			log.Errorw(v.Name(), "spec", v.Spec(), "err", err, "entiryId", entryId)
			return err
		} else {
			n.jobEntryId = append(n.jobEntryId, entryId)
		}
	}
	n.cron.Start()
	return nil
}

func (n StartJob) Stop() context.Context {
	return n.cron.Stop()
}

func warpJob(name string, fn JobHandler) cron.Job {
	ctx := context.TODO()
	return cron.NewChain(cron.SkipIfStillRunning(newLogger(ctx))).Then(
		cron.FuncJob(
			func() {
				//ctx, span := tracing.StartConsume(context.Background(), name)
				//span.End()
				defer func() {
					if p := recover(); p != nil {
						stack := Stack(3)
						log.Errorf("脚本执行异常: %s %s", p, stack)
					}
				}()

				begin := time.Now()
				log.Infof("开始执行脚本: %s", name)
				fn(ctx)
				log.Infof("结束执行脚本: %s, 耗时%.3fs", name, time.Since(begin).Seconds())
			},
		),
	)
}
