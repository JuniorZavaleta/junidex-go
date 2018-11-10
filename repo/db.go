package repo

import (
	"database/sql"
	"github.com/go-redis/redis"
	"os"
)

var db *sql.DB

var client *redis.Client

func InitDatabase() {
	var err error

	db, err = sql.Open("mysql", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})
}
