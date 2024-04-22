package telegram

import (
	"context"
	"fmt"
	"net/http"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service struct {
	bot *tgbotapi.BotAPI
}

var _ peacefulroad.TelegramService = (*Service)(nil)

func New(token string) *Service {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic("cannot create tg bot: " + err.Error())
	}
	return &Service{
		bot: bot,
	}
}

func (s *Service) GetBotUrl() string {
	return fmt.Sprintf("https://t.me/%s", s.bot.Self.UserName)
}

type UpdateHandler func(context.Context, tgbotapi.Update) error

func (s *Service) HandleUpdatesPull(handler UpdateHandler) {
	updatesChannel := s.bot.GetUpdatesChan(tgbotapi.UpdateConfig{})
	for {
		upd := <-updatesChannel
		go func() {
			err := handler(context.Background(), upd)
			if err != nil {
				fmt.Println("error", err)
			}
		}()
	}
}

func (s *Service) HandleUpdatesEndpoint(handler UpdateHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upd, err := s.bot.HandleUpdate(r)
		if err != nil {
			http.Error(w, "cannot parse update", http.StatusBadRequest)
			return
		}

		err = handler(r.Context(), *upd)

		if err != nil {
			http.Error(w, "cannot handle update", http.StatusInternalServerError)
		}
	})
}
