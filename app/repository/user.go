package repository

import (
	"challange/app/infrastracture"
	"challange/app/interfaces"
	"challange/app/models"
	"context"
	"time"
)

type UserRepository struct {
	logger   interfaces.Logger
	db       interfaces.DB
	memoryDb interfaces.MemoryDB
}

func NewUserRepository(
	logger infrastracture.SegmentLogger,
	db infrastracture.PgxDB,
	memoryDB infrastracture.Redis) UserRepository {
	return UserRepository{
		db:       &db,
		logger:   &logger,
		memoryDb: &memoryDB,
	}
}

func (us *UserRepository) Save(userID string, segment string, ExpiredSegment time.Time) error {
	ctx := context.TODO()
	parameters := []interface{}{
		userID, segment, ExpiredSegment,
	}
	_, err := us.db.Exec(
		ctx,
		"INSERT INTO users (ID,segment,expired_segment) values($1,$2,$3)",
		parameters)
	return err
}

func (us *UserRepository) List() (users []*models.User, err error) {
	ctx := context.TODO()

	values, err := us.db.Query(
		ctx,
		"SELECT * FROM users",
		[]interface{}{})
	if err != nil {
		return
	}
	for _, v := range values {
		users = append(users, &models.User{
			ID:            v[0].(string),
			Segment:       v[1].(string),
			ExpireSegment: v[2].(time.Time),
		})
	}
	return
}
