package auth

import (
	"fmt"
	"net/http"
)

// Handler for our login.
func (a *Auth) logoutHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	chatId, err := getChatId(req)
	if err != nil {
		err = fmt.Errorf("invalid query parameters: %w", err)
		a.log.WarnContext(ctx, err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := a.us.GetUser(ctx, chatId)
	if err != nil {
		err = fmt.Errorf("cannot find user: %w", err)
		a.log.WarnContext(ctx, err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.Token = nil

	err = a.us.SaveUser(ctx, u)
	if err != nil {
		err = fmt.Errorf("cannot remove user session: %w", err)
		a.log.ErrorContext(ctx, err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// err = a.tgs.SendWelcomeMessage(ctx, u)
	// if err != nil {
	// 	err = fmt.Errorf("cannot remove user session: %w", err)
	// 	a.log.ErrorContext(ctx, err.Error())
	// 	return
	// }

	http.Redirect(w, req, a.tgs.GetBotUrl(), http.StatusTemporaryRedirect)
}
