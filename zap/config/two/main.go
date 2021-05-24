package main

import (
	"go-daily-lib/zap/config/util"
	"time"
)

func main() {
	l := util.GetLogger()
	for i := 1; i <= 20_000; i++ {
		time.Sleep(time.Millisecond)
		l.Info("2")
	}
}
