package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	curTime := time.Now()
	fiveMonthLater := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+70, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.Local)
	diff := fiveMonthLater.Sub(curTime).Hours()
	logrus.Info("time diff: ", diff)
}
