package tests

import (
	"allergycron/utils"
	"testing"
)

func TestSendSlackMessage(t *testing.T) {
	err := utils.SendSlackMessage("Test message!")
	if err != nil {
		t.Errorf("Error sending Slack message: %s", err)
	}
}
