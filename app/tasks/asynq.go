package tasks

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"github.com/hibiken/asynq"
	"os"
)

//TaskAsynq -> TaskAsynq Struct
type TaskAsynq struct {
	Logger interfaces.Logger
	Server *asynq.Server
}

//NewTaskAsynq -> return new TaskAsynq struct,
func NewTaskAsynq(
	logger infrastracture.SegmentLogger,
) TaskAsynq {
	return TaskAsynq{
		Logger: &logger,
		Server: asynq.NewServer(asynq.RedisClientOpt{Addr: os.Getenv("")},
			asynq.Config{
				Concurrency: 10,
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"info":     1,
				},
			},
		),
	}
}

//NewClient -> return asynq client don't forget to close it
func (t *TaskAsynq) NewClient() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("RedisAddr")})
}
