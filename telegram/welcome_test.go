package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testChatId = 108034197

func TestWelcomeMessage(t *testing.T) {
	tg := New(token)

	err := tg.SendWelcomeMessage(testChatId, "http://pacebot.test:8080/login?chatId=108034197")
	assert.NoError(t, err)
}
