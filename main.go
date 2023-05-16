package main

import (
	"time"

	"github.com/robfig/cron"
)

func main() {
	loc, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	cronJob := cron.NewWithLocation(loc)
	cronJob.AddFunc("0 0 8 * * *", func() {
	})
}
