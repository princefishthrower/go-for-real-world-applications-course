package main

import (
	"allergycron/allergy_api"
	"allergycron/utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loc, err := time.LoadLocation(os.Getenv("CRON_TIMEZONE"))
	if err != nil {
		panic(err)
	}

	// print CRON_SCHEDULE to terminal
	fmt.Println("CRON_SCHEDULE: " + os.Getenv("CRON_SCHEDULE"))

	cronJob := cron.NewWithLocation(loc)
	cronJob.AddFunc(os.Getenv("CRON_SCHEDULE"), func() {
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
