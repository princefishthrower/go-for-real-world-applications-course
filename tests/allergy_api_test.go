package tests

import (
	"allergycron/allergy_api"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestAllergyApi(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	message, err := allergy_api.GetHourlyLoadData()
	if err != nil {
		t.Errorf("Error getting hourly load data: %s", err)
	}
	if message == nil {
		t.Errorf("Error getting hourly load data: message is nil")
	}
	if *message == "" {
		t.Errorf("Error getting hourly load data: message is empty")
	}
	message, err = allergy_api.GetCurrentChartData()
	if err != nil {
		t.Errorf("Error getting current chart data: %s", err)
	}
	if message == nil {
		t.Errorf("Error getting current chart data: message is nil")
	}
	if *message == "" {
		t.Errorf("Error getting current chart data: message is empty")
	}
}
