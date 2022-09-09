package main

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	iter := rdb.Scan(0, "hello*", 0).Iterator()
	for iter.Next() {
		logrus.New().Info("key: ", iter.Val())
		// if err != nil { panic(err) }
	}
}
