package logic

import (
	"github.com/JakubC-projects/pacebot"
	"github.com/JakubC-projects/pacebot/auth"
	"github.com/JakubC-projects/pacebot/myshare"
)

type Logic struct {
	tg             pacebot.TelegramService
	us             pacebot.UserService
	ms             *myshare.Client
	auth           *auth.Auth
	allowedClubIds []int
}

func New(tg pacebot.TelegramService, us pacebot.UserService, ms *myshare.Client, auth *auth.Auth, clubIds []int) *Logic {
	l := &Logic{tg, us, ms, auth, clubIds}
	auth.SetPostLoginAction(l.postLoginHook)
	auth.SetPostLogoutAction(l.postLogoutHook)
	return l
}
