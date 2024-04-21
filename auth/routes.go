// platform/router/router.go

package auth

import (
	"fmt"
	"net/http"
)

func (a *Auth) AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /login", a.loginHandler)
	mux.HandleFunc("GET /callback", a.callbackHandler)
	mux.HandleFunc("GET /logout", a.logoutHandler)
}

func (a *Auth) LoginEndpoint() string {
	return fmt.Sprintf("%s/login", a.host)
}
