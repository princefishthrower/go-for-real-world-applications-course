package allergy_api

import (
	"allergycron/utils"
	"fmt"
	"math"
	"net/url"
	"os"
	"time"
)

type HourlyLoadResult struct {
	Total  int   `json:"total"`
	Hourly []int `json:"hourly"`
}

type HourlyLoadResponse struct {
	Success int              `json:"success"`
	Result  HourlyLoadResult `json:"result"`
}

type CurrentChartDataResult struct {
	Date    string  `json:"date"`
	Average float64 `json:"average"`
}

type CurrentChartDataResponse struct {
	Success int                      `json:"success"`
	Results []CurrentChartDataResult `json:"results"`
}

func GetHourlyLoadData() (*string, error) {
	queryParameters := url.Values{}
	queryParameters.Add("eID", "appinterface")
	queryParameters.Add("action", "getHourlyLoadData")
	queryParameters.Add("type", "zip")
	queryParameters.Add("zip", "6800")
	queryParameters.Add("country", "AT")
	queryParameters.Add("pure_json", "1")

	response, err := utils.MakeHTTPRequest(os.Getenv("ALLERGY_API_URL_ROOT"), "GET", nil, queryParameters, nil, HourlyLoadResponse{})
	if err != nil {
		return nil, err
	}

	averageLoad := 0
	for _, hour := range response.Result.Hourly {
		averageLoad += hour
	}
	averageLoad = averageLoad / len(response.Result.Hourly)

	scaledAverageLoad := averageLoad / 2

	formattedMessage := formatAllergyData(scaledAverageLoad)

	return &formattedMessage, nil
}

func GetCurrentChartData() (*string, error) {
	queryParameters := url.Values{}
	queryParameters.Add("eID", "appinterface")
	queryParameters.Add("action", "getCurrentChartData")
	queryParameters.Add("poll_id", "5")
	queryParameters.Add("zip", "6800")
	queryParameters.Add("season", "2")
	queryParameters.Add("pure_json", "1")

	response, err := utils.MakeHTTPRequest(os.Getenv("ALLERGY_API_URL_ROOT"), "GET", nil, queryParameters, nil, CurrentChartDataResponse{})
	if err != nil {
		return nil, err
	}

	currentYYYYMMDD := time.Now().Format("2006-01-02")
	averageHistorical := 0.0
	for _, result := range response.Results {
		if result.Date == currentYYYYMMDD {
			averageHistorical = result.Average
		}
	}

	scaledAverageHistorical := int(math.Round(averageHistorical / 2.0))

	formattedMessage := fmt.Sprintf("Historically, the average pollen load for today is %d", scaledAverageHistorical)

	return &formattedMessage, nil
}

func formatAllergyData(scaledAverageLoad int) string {
	formattedMessage := fmt.Sprintf("The average pollen load for today is %d", scaledAverageLoad)

	switch {
	case scaledAverageLoad == 1:
		return "🟡 Okay. " + formattedMessage + " (LOW). 🟡"
	case scaledAverageLoad == 2:
		return "🟠 Watch out! " + formattedMessage + " (MEDIUM)! 🟠"
	case scaledAverageLoad == 3:
		return "🔴 Warning! " + formattedMessage + " (HIGH)! 🔴"
	case scaledAverageLoad == 4:
		return "🔴🔴🔴 Alert! " + formattedMessage + " (VERY HIGH)! 🔴🔴🔴"
	case scaledAverageLoad == 0:
		fallthrough
	default:
		return "🟢 Nice! " + formattedMessage + " (NONE). 🟢"
	}

}
