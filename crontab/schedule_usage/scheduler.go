package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {

	fmt.Println("Starting...")

	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("Tick every 1 second")
	})

	c.AddFunc("@every 1hour", func() {
		fmt.Println("Tick every hours")
	})

	c.AddFunc("@every 3s", func() {
		fmt.Println("Tick every 3 seconds")
	})

	c.AddFunc("@every 5s", func() {
		fmt.Println("Tick every 5 seconds")
	})

	c.AddFunc("@daily", func() {
		fmt.Println("Every day on midnight")
	})

	c.AddFunc("@weekly", func() {
		fmt.Println("Every week")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
