package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func ClearCacheWithPrefix(prefix string) {
	prefix = prefix + "*"
	iter := rdb.Scan(0, prefix, 0).Iterator()
	for iter.Next() {
		err := rdb.Del(iter.Val()).Err()
		if err != nil {
			logrus.Errorf("error when delete key %s", iter.Val())
			return
		}
	}
}
