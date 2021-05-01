package models

import (
	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}