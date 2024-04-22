package auth

import (
	"errors"
	"fmt"
	"net/http"

	peacefulroad "github.com/JakubC-projects/peaceful-road"
	"github.com/lestrrat-go/jwx/jwt"
	"golang.org/x/oauth2"
)

func (a *Auth) callbackHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ls, ok := getLoginCookie(req)
	if !ok {
		err := fmt.Errorf("cannot find login session")
		a.log.WarnContext(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.getCallbackToken(req, ls)
	if err != nil {
		err := fmt.Errorf("cannot get callback token: %w", err)
		a.log.WarnContext(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)

		return
	}

	user := peacefulroad.User{ChatId: ls.ChatId}
	err = fillUserDataFromToken(&user, token)

	if err != nil {
		err := fmt.Errorf("cannot parse token: %w", err)
		a.log.ErrorContext(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if a.postLoginAction != nil {
		err = a.postLoginAction(ctx, user)
		if err != nil {
			err := fmt.Errorf("cannot perform post login action: %w", err)
			a.log.ErrorContext(ctx, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	http.Redirect(w, req, a.tgs.GetBotUrl(), http.StatusTemporaryRedirect)
}

func (a *Auth) getCallbackToken(req *http.Request, ls loginState) (*oauth2.Token, error) {

	if req.URL.Query().Get("state") != ls.State {
		return nil, errors.New("invalid state parameter")
	}

	token, err := a.config.Exchange(req.Context(), req.URL.Query().Get("code"))
	if err != nil {
		return nil, errors.New("failed to exchange an authorization code for a token")
	}
	return token, nil
}

func fillUserDataFromToken(user *peacefulroad.User, token *oauth2.Token) error {
	idToken := token.Extra("id_token").(string)
	idTokenClaims, err := jwt.Parse([]byte(idToken))
	if err != nil {
		return fmt.Errorf("cannot parse id token claims: %w", err)
	}

	userName, ok := idTokenClaims.Get("name")
	if !ok {
		return errors.New("cannot find name claim")
	}
	userNameString, ok := userName.(string)
	if !ok {
		return errors.New("invalid type of name claim ")
	}

	personId, ok := idTokenClaims.Get("https://login.bcc.no/claims/personId")
	if !ok {
		return errors.New("cannot find personId claim")
	}
	personIdFloat, ok := personId.(float64)
	if !ok {
		return errors.New("invalid type of personId claim ")
	}

	user.Token = token
	user.DisplayName = userNameString
	user.PersonID = int(personIdFloat)

	return nil
}
