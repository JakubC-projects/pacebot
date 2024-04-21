package peacefulroad

type TelegramService interface {
	GetBotUrl() string
	SendWelcomeMessage(chatId int, loginUrl string) error
	SendStatusMessage(chatId int) error
	// SendStatusMessage(ctx context.Context, u User, status Status) error
}
