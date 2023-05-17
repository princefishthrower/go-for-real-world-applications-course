package main

import (
	"allergycron/allergy_api"
	"allergycron/utils"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	loc, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	cronJob := cron.NewWithLocation(loc)
	cronJob.AddFunc("0 15 12 * * *", func() {
		dailyAverageMessage, err := allergy_api.GetHourlyLoadData()
		if err != nil {
			panic(err)
		}

		historicalAverageMessage, err := allergy_api.GetCurrentChartData()
		if err != nil {
			panic(err)
		}

		slackMessage := *dailyAverageMessage + "\n" + *historicalAverageMessage
		err = utils.SendSlackMessage(slackMessage)
		if err != nil {
			panic(err)
		}

		log.Println("Successfully sent Slack message: " + slackMessage)
	})

	cronJob.Start()

	select {}
}
