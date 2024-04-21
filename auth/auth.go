package auth

import (
	"context"
	"fmt"
	"log/slog"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type PostLoginAction func(ctx context.Context, user peacefulroad.User) error

type Auth struct {
	config          *oauth2.Config
	host            string
	us              peacefulroad.UserService
	tgs             peacefulroad.TelegramService
	log             *slog.Logger
	postLoginAction PostLoginAction
}

type Config struct {
	Issuer, ClientId, ClientSecret, Audience, Host string
}

func New(
	conf Config,
	us peacefulroad.UserService,
	tgs peacefulroad.TelegramService,
	log *slog.Logger,
) *Auth {

	endpoints := oauth2.Endpoint{
		AuthURL:  fmt.Sprintf("https://%s/oauth/authorize?audience=%s", conf.Issuer, conf.Audience),
		TokenURL: fmt.Sprintf("https://%s/oauth/token", conf.Issuer),
	}

	oauthConfig := &oauth2.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s/callback", conf.Host),
		Endpoint:     endpoints,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "offline_access"},
	}

	return &Auth{oauthConfig, conf.Host, us, tgs, log, nil}
}
