package auth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

type loginState struct {
	State  string
	ChatId int
}

const cookieName = "login-state"

func addLoginCookie(w http.ResponseWriter, ls loginState) {
	stateJson, _ := json.Marshal(ls)
	stateBase64 := base64.StdEncoding.EncodeToString(stateJson)

	cookie := http.Cookie{
		Name:     cookieName,
		Value:    stateBase64,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(time.Hour)}

	http.SetCookie(w, &cookie)
}

func getLoginCookie(req *http.Request) (loginState, bool) {
	stateBase64, err := req.Cookie(cookieName)
	if err != nil {
		return loginState{}, false
	}
	stateJson, err := base64.StdEncoding.DecodeString(stateBase64.Value)
	if err != nil {
		return loginState{}, false
	}

	var state loginState
	err = json.Unmarshal(stateJson, &state)
	if err != nil {
		return state, false
	}
	return state, true
}
