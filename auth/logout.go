package auth

import (
	"fmt"
	"net/http"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
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

	err = a.postLogoutAction(ctx, peacefulroad.User{ChatId: chatId})
	if err != nil {
		err := fmt.Errorf("cannot perform post logout action: %w", err)
		a.log.ErrorContext(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, a.logoutUrl, http.StatusTemporaryRedirect)
}
