package redis

import (
	"errors"
	"playground/constants"
	"time"

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

func LockWithKey(key string) error {
	res := rdb.SetNX(key, 1, constants.RedisLockExpiredTime*time.Second)
	lockSuccess, err := res.Result()

	if err != nil || !lockSuccess {
		return errors.New("can not get lock")
	}

	return nil
}

func ReleaseLockWithKey(key string) error {
	res := rdb.Del(key)

	unlockSuccess, err := res.Result()
	if err == nil && unlockSuccess > 0 {
		return nil
	} else {
		return errors.New("unlock failed")
	}
}
