package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/JakubC-projects/peaceful-road/auth"
	"github.com/JakubC-projects/peaceful-road/inmem"
	"github.com/JakubC-projects/peaceful-road/logic"
	"github.com/JakubC-projects/peaceful-road/telegram"
)

var (
	port       = os.Getenv("PORT")
	serverHost = os.Getenv("SERVER_HOST")

	telegramApiKey = os.Getenv("TELEGRAM_API_KEY")

	oauthIssuer       = os.Getenv("OAUTH_ISSUER")
	oauthClientId     = os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")

	// myshareBaseUrl  = os.Getenv("MYSHARE_BASE_URL")
	myshareAudience = os.Getenv("MYSHARE_AUDIENCE")
)

func main() {
	mux := http.NewServeMux()

	us := inmem.NewUserService()
	tg := telegram.New(telegramApiKey)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))

	auth := auth.New(auth.Config{
		Issuer:       oauthIssuer,
		ClientId:     oauthClientId,
		ClientSecret: oauthClientSecret,
		Audience:     myshareAudience,
		Host:         serverHost,
	}, us, tg, logger)

	logic := logic.New(tg, us, auth)

	auth.AddRoutes(mux)

	go tg.HandleUpdatesPull(logic.HandleUpdate)

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
