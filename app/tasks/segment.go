package tasks

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeCountSegmentUser = "count:segmentUsers"
)

type SegmentTask struct {
	logger interfaces.Logger
}

func NewSegmentTask(logger infrastracture.SegmentLogger) SegmentTask {
	return SegmentTask{
		logger: &logger,
	}
}

func (et *SegmentTask) NewCountSegmentTask() (*asynq.Task, error) {
	return asynq.NewTask(
		TypeCountSegmentUser,
		[]byte{},
		asynq.Timeout(80*time.Second),
		asynq.MaxRetry(2)), nil
}

func (et SegmentTask) HandleCountSegmentTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println("handling verify count segment")
	return nil
}
