package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/JakubC-projects/pacebot/auth"
	"github.com/JakubC-projects/pacebot/firebase"
	"github.com/JakubC-projects/pacebot/logic"
	"github.com/JakubC-projects/pacebot/myshare"
	"github.com/JakubC-projects/pacebot/telegram"
	"github.com/samber/lo"
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
	myshareClubIds  = os.Getenv("MYSHARE_CLUB_IDS")
	myshareAudience = os.Getenv("MYSHARE_AUDIENCE")

	usePolling = os.Getenv("USE_POLLING")
)

func main() {
	mux := http.NewServeMux()

	tg := telegram.New(telegramApiKey)

	myshareClient := myshare.NewClient(myshareBaseUrl)
	firestore := firebase.NewStore(gcpProjectId)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))

	auth := auth.New(auth.Config{
		Issuer:       oauthIssuer,
		ClientId:     oauthClientId,
		ClientSecret: oauthClientSecret,
		Audience:     myshareAudience,
		Host:         serverHost,
	}, firestore, tg, logger)

	logic := logic.New(tg, firestore, myshareClient, auth, parseClubIds())

	auth.AddRoutes(mux)

	if usePolling == "true" {
		go tg.HandleUpdatesPull(logic.HandleUpdate)
	} else {
		mux.Handle("/updates", tg.HandleUpdatesEndpoint(logic.HandleUpdate))
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func parseClubIds() []int {
	clubIdsRaw := strings.Split(myshareClubIds, ",")
	return lo.Map(clubIdsRaw, func(idRaw string, _ int) int {
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic("cannot parse club ids: " + err.Error())
		}
		return id
	})
}
