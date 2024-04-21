package telegram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusMessage(t *testing.T) {
	tg := New(token)

	err := tg.SendStatusMessage(testChatId)
	assert.NoError(t, err)
}
