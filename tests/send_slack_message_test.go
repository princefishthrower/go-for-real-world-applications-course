package tests

import (
	"allergycron/utils"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendSlackMessage(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = utils.SendSlackMessage("Test message!")
	if err != nil {
		t.Errorf("Error sending Slack message: %s", err)
	}
}
