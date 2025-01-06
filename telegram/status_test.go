package telegram

import (
	"testing"

	"github.com/JakubC-projects/pacebot"
	"github.com/stretchr/testify/assert"
)

func TestStatusMessage(t *testing.T) {
	tg := New(token)

	s := pacebot.StatusMessage{
		CurrentStatus:   1000,
		SeasonTarget:    2000,
		MilestoneTarget: 40,
		Currency:        "PLN",

		RegisterURL: "https://app.myshare.today/registration",
		DonateURL:   "https://donationbuk.no",

		LogoutURL: "http://pacebot.test:8080/logout?chatId=108034197",
	}

	err := tg.SendStatusMessage(testChatId, s)
	assert.NoError(t, err)
}
