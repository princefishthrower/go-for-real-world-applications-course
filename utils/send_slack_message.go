package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func SendSlackMessage(message string) error {
	body, err := json.Marshal(map[string]string{"text": message})
	if err != nil {
		return err
	}

	// log SLACK_WEBHOOK_URL to stdout
	fmt.Println(os.Getenv("SLACK_WEBHOOK_URL"))

	MakeHTTPRequest(os.Getenv("SLACK_WEBHOOK_URL"), "POST", nil, nil, bytes.NewBuffer(body), "")

	return nil
}
