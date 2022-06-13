package tests

import (
	"challange/app/infrastracture"
	"context"
	"testing"
	"time"
)

func TestPgxExec(t *testing.T) {
	l := infrastracture.NewLogger()
	db := infrastracture.NewPgxDB(l)
	rand := infrastracture.NewRandom()
	rand.RefreshSeed()
	id := rand.GenerateRandomStr(10)
	segment := rand.GenerateRandomStr(10)
	expired := time.Now().Add(time.Hour * time.Duration(rand.RandomNumber(700)))
	rowsAffected, err := db.Exec(
		context.Background(),
		"INSERT INTO USERS (ID,segment,expired_segment) values($1,$2,$3) ",
		[]interface{}{id, segment, expired},
	)
	if rowsAffected < 1 || err != nil {
		t.Errorf("Pgx %d rows effected: %s", rowsAffected, err.Error())
	}
}
