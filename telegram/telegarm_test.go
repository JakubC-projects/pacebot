package telegram

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token = os.Getenv("TELEGRAM_API_KEY")

func TestPullUpdates(t *testing.T) {
	tg := New(token)

	wg := sync.WaitGroup{}
	wg.Add(1)

	tg.HandleUpdatesPull(func(ctx context.Context, u tgbotapi.Update) error {
		fmt.Println(u)
		wg.Done()

		return nil
	})

	wg.Wait()
}
