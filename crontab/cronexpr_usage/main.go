package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	// Linux crontab
	// Year configuration (2021-2099)

	// Every 5 seconds execute
	if expr, err = cronexpr.Parse("*/5 * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 0, 6, 12, 18, ...

	// Current time
	now = time.Now()

	// Next call time
	nextTime = expr.Next(now)

	// Waiting timer timeout
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("Has been called:", nextTime)
	})

	time.Sleep(5 * time.Second)

}
