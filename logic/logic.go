package logic

import (
	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"github.com/JakubC-projects/peaceful-road/auth"
	"github.com/JakubC-projects/peaceful-road/myshare"
)

type Logic struct {
	tg   peacefulroad.TelegramService
	us   peacefulroad.UserService
	ms   *myshare.Client
	auth *auth.Auth
}

func New(tg peacefulroad.TelegramService, us peacefulroad.UserService, ms *myshare.Client, auth *auth.Auth) *Logic {
	l := &Logic{tg, us, ms, auth}
	auth.SetPostLoginAction(l.postLoginHook)
	auth.SetPostLogoutAction(l.postLogoutHook)
	return l
}
