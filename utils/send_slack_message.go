package utils

import (
	"bytes"
	"encoding/json"
	"os"
)

func SendSlackMessage(message string) error {
	body, err := json.Marshal(map[string]string{"text": message})
	if err != nil {
		return err
	}

	MakeHTTPRequest(os.Getenv("SLACK_WEBHOOK_URL"), "POST", nil, nil, bytes.NewBuffer(body), "")

	return nil
}
