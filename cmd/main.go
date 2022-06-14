package main

import (
	"challange/app/infrastracture"
	"context"
	"log"
	"os"
	"time"
)

//send console argument for executing command like "go run ./cmd/ seed"
func main() {
	arg := os.Args[len(os.Args)-1]
	switch arg {
	//create random amount of users
	case "seed":
		rand := infrastracture.NewRandom()
		logger := infrastracture.NewLogger()
		db := infrastracture.NewPgxDB(logger)
		rand.RefreshSeed()
		for i := 0; i < rand.RandomNumber(1000); i++ {
			rand.RefreshSeed()
			id := rand.GenerateRandomStr(10)
			segment := rand.GenerateRandomStr(2)
			expired := time.Now().Add(time.Hour * time.Duration(rand.RandomNumber(700)))
			parameters := []interface{}{
				id, segment, expired,
			}
			_, err := db.Exec(context.TODO(), "INSERT INTO users values($1,$2,$3)", parameters)
			if err != nil {
				logger.Error(err.Error())
			}
		}
	default:
		log.Fatal("Unkown command!")
	}
}
