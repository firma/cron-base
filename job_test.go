package job

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	job := NewJob()
	job.RegisterJob(NewJobInfo("test", "* * * * * *", hnadler))
	if err := job.Start(); err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := job.Stop(); err != nil {
		log.Println(err)
	}
	log.Println("job exiting")

	select {
	case <-ctx.Done():
		log.Println("超时结束")
		break
	}

	log.Println("退出")
}

func hnadler(ctx context.Context) error {
	log.Println("xxx")
	source := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(source)
	i := ran.Int()
	if i%2 == 0 {
		panic("xxxxx")
	}

	return nil
}
