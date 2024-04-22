package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
)

// Handler for our login.
func (a *Auth) loginHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	state, err := generateRandomState()
	if err != nil {
		errCtx := fmt.Errorf("cannot generate random state: %w", err)

		a.log.WarnContext(ctx, errCtx.Error())
		http.Error(w, errCtx.Error(), http.StatusInternalServerError)
		return
	}

	chatId, err := getChatId(req)
	if err != nil {
		errCtx := fmt.Errorf("invalid query parameters: %w", err)

		a.log.WarnContext(ctx, errCtx.Error())
		http.Error(w, errCtx.Error(), http.StatusBadRequest)
		return
	}

	addLoginCookie(w, loginState{
		State:  state,
		ChatId: chatId,
	})

	http.Redirect(w, req, a.config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func getChatId(req *http.Request) (int, error) {
	chatIdRaw := req.URL.Query().Get("chatId")
	if chatIdRaw == "" {
		return 0, fmt.Errorf("chatId is missing")
	}
	chatId, err := strconv.Atoi(chatIdRaw)
	return chatId, err
}

func (a *Auth) SetPostLoginAction(action PostLoginAction) {
	a.postLoginAction = action
}

func (a *Auth) SetPostLogoutAction(action PostLoginAction) {
	a.postLogoutAction = action
}
