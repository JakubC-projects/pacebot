package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type PostLoginAction func(ctx context.Context, user peacefulroad.User) error

type Auth struct {
	config           *oauth2.Config
	logoutUrl        string
	host             string
	tgs              peacefulroad.TelegramService
	log              *slog.Logger
	postLoginAction  PostLoginAction
	postLogoutAction PostLoginAction
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

	logoutUrl := fmt.Sprintf("https://%s/v2/logout?returnTo=%s&client_id=%s ", conf.Issuer, url.QueryEscape(tgs.GetBotUrl()), conf.ClientId)

	oauthConfig := &oauth2.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  fmt.Sprintf("%s/callback", conf.Host),
		Endpoint:     endpoints,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "offline_access"},
	}

	return &Auth{oauthConfig, logoutUrl, conf.Host, tgs, log, nil, nil}
}
