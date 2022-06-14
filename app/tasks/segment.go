package tasks

import (
	"bytes"
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/models"
	"context"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeCountSegmentUser = "count:segmentUsers"
)

type SegmentTask struct {
	logger   interfaces.Logger
	memoryDB interfaces.MemoryDB
	db       interfaces.DB
}

func NewSegmentTask(
	logger infrastracture.SegmentLogger,
	redis infrastracture.Redis,
	db infrastracture.PgxDB) SegmentTask {
	return SegmentTask{
		logger:   &logger,
		memoryDB: &redis,
		db:       &db,
	}
}

func (et *SegmentTask) NewCountSegmentTask() (*asynq.Task, error) {
	return asynq.NewTask(
		TypeCountSegmentUser,
		[]byte{},
		asynq.Timeout(80*time.Second),
		asynq.MaxRetry(2)), nil
}

//this method get count of segments users and Store count of segment users in memory db
func (et SegmentTask) HandleCountSegmentTask(ctx context.Context, t *asynq.Task) error {
	values, err := et.db.Query(ctx,
		"SELECT segment,COUNT(segment) FROM users GROUP BY segment",
		[]interface{}{})
	if err != nil {
		et.logger.Error("error in handling task" + err.Error())
	}
	var buff bytes.Buffer
	for i, value := range values {
		s := &models.Segment{
			Title:      value[0].(string),
			UsersCount: value[1].(int64),
		}
		if i > 0 {
			buff.Write([]byte(","))
		}
		s.ToJson(&buff)
	}
	jsonStr := "{" + buff.String() + "}"

	et.memoryDB.Set("segments", jsonStr, 24*time.Hour)
	return nil
}
