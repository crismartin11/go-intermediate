package handlers

import (
	"net/http"

	"go-intermediate/pkg/okta"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	// Invoco a login de Okta
	var redirectPath string = okta.GetPathOkta(request)
	http.Redirect(response, request, redirectPath, http.StatusFound)
}

func AuthCodeCallbackHandler(response http.ResponseWriter, request *http.Request) {
	// Luego de autenticación correcta, retorna a este endpoint
	okta.SaveSession(response, request)
	http.Redirect(response, request, "/", http.StatusFound)
}

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	okta.RemoveSession(response, request)
	http.Redirect(response, request, "/", http.StatusFound)
}
