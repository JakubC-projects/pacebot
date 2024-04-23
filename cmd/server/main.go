package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/JakubC-projects/peaceful-road/auth"
	"github.com/JakubC-projects/peaceful-road/firebase"
	"github.com/JakubC-projects/peaceful-road/logic"
	"github.com/JakubC-projects/peaceful-road/myshare"
	"github.com/JakubC-projects/peaceful-road/telegram"
)

var (
	port       = os.Getenv("PORT")
	serverHost = os.Getenv("SERVER_HOST")

	telegramApiKey = os.Getenv("TELEGRAM_API_KEY")
	gcpProjectId   = os.Getenv("GCP_PROJECT")

	oauthIssuer       = os.Getenv("OAUTH_ISSUER")
	oauthClientId     = os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret = os.Getenv("OAUTH_CLIENT_SECRET")

	myshareBaseUrl  = os.Getenv("MYSHARE_BASE_URL")
	myshareClubId   = os.Getenv("MYSHARE_CLUB_ID")
	myshareAudience = os.Getenv("MYSHARE_AUDIENCE")

	usePolling = os.Getenv("USE_POLLING")
)

func main() {
	os.Setenv("TZ", "Europe/Warsaw")

	mux := http.NewServeMux()

	tg := telegram.New(telegramApiKey)

	myshareClient := myshare.NewClient(myshareBaseUrl, myshareClubId)
	firestore := firebase.NewStore(gcpProjectId)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))

	auth := auth.New(auth.Config{
		Issuer:       oauthIssuer,
		ClientId:     oauthClientId,
		ClientSecret: oauthClientSecret,
		Audience:     myshareAudience,
		Host:         serverHost,
	}, firestore, tg, logger)

	logic := logic.New(tg, firestore, myshareClient, auth)

	auth.AddRoutes(mux)

	if usePolling == "true" {
		go tg.HandleUpdatesPull(logic.HandleUpdate)
	} else {
		mux.Handle("/updates", tg.HandleUpdatesEndpoint(logic.HandleUpdate))
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
