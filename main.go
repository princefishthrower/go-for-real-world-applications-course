package main

import (
	"allergycron/allergy_api"
	"allergycron/utils"
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

	// log CRON_SCHEDULE to stdout
	log.Println(os.Getenv("CRON_SCHEDULE"))

	// log entire env to stdout
	log.Println(os.Environ())
	log.Println(os.Environ())

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
