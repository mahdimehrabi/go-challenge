package tasks

import (
	"challange/app/interfaces"
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewTaskAsynq),
	fx.Provide(NewSegmentTask),
)

type Task interface {
	HandlesToMux() error
}

type Tasks struct {
	logger      interfaces.Logger
	taskAsynq   TaskAsynq
	segmentTask SegmentTask
}

func NewTasks(
	logger interfaces.Logger,
	taskAsynq TaskAsynq,
	segmentTask SegmentTask) Tasks {
	return Tasks{
		logger:      logger,
		taskAsynq:   taskAsynq,
		segmentTask: segmentTask,
	}
}

func (t *Tasks) HandleTasks() error {
	serverMux := asynq.NewServeMux()
	serverMux.HandleFunc(
		TypeCountSegmentUser,
		t.segmentTask.HandleCountSegmentTask,
	)
	return t.taskAsynq.Server.Run(serverMux)
}
