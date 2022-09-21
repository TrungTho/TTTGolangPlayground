package main

import (
	"playground/redis"
)

type RewardServiceError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func main() {
	redis.ClearCacheWithPrefix("hello")
}
