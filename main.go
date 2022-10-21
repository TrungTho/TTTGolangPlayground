package main

import (
	"playground/cryptos"
)

type RewardServiceError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func main() {
	cryptos.Demo()
}
