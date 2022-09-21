package main

import (
	"playground/handlers"
)

type RewardServiceError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func main() {
	handlers.Setup()
}
