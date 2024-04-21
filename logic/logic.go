package logic

import (
	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"github.com/JakubC-projects/peaceful-road/auth"
)

type Logic struct {
	tg       peacefulroad.TelegramService
	us       peacefulroad.UserService
	loginUrl string
}

func New(tg peacefulroad.TelegramService, us peacefulroad.UserService, auth *auth.Auth) *Logic {
	l := &Logic{tg, us, auth.LoginEndpoint()}
	auth.SetPostLoginAction(l.postLoginHook)
	return l
}
