package peacefulroad

type StatusMessage struct {
	CurrentStatus float64
	SeasonTarget  float64

	Currency string

	WeekTarget float64

	RegisterURL string
	DonateURL   string

	LogoutURL string
}

type TelegramService interface {
	GetBotUrl() string
	SendWelcomeMessage(chatId int, loginUrl string) error
	SendStatusMessage(chatId int, status StatusMessage) error
	EditStatusMessage(chatId int, messageId int, status StatusMessage) error
}
