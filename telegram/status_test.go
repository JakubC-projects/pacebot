package telegram

import (
	"testing"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"github.com/stretchr/testify/assert"
)

func TestStatusMessage(t *testing.T) {
	tg := New(token)

	s := peacefulroad.StatusMessage{
		CurrentStatus:   1000,
		SeasonTarget:    2000,
		MilestoneTarget: 40,
		Currency:        "PLN",

		RegisterURL: "https://app.myshare.today/registration",
		DonateURL:   "https://donationbuk.no",

		LogoutURL: "http://peacefulroad.test:8080/logout?chatId=108034197",
	}

	err := tg.SendStatusMessage(testChatId, s)
	assert.NoError(t, err)
}
