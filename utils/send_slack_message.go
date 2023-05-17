package utils

import (
	"bytes"
	"encoding/json"
)

func SendSlackMessage(message string) error {
	body, err := json.Marshal(map[string]string{"text": message})
	if err != nil {
		return err
	}

	MakeHTTPRequest("https://hooks.slack.com/services/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", "POST", nil, nil, bytes.NewBuffer(body), "")

	return nil
}
