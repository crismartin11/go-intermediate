package okta

import (
	"fmt"
	"net/http"
	"os"

	"go-intermediate/pkg/okta/utils"

	"github.com/gorilla/sessions"
)

const SESSION_KEY string = "okta-hosted-login-session-store"

var (
	state        = utils.GenerateState()
	sessionStore = sessions.NewCookieStore([]byte(SESSION_KEY))
)

func GetPathOkta(request *http.Request) string {
	return os.Getenv("ISSUER") + "/v1/authorize?" + utils.GetQueryLogin(request, state)
}

func SaveSession(response http.ResponseWriter, request *http.Request) {
	// Me retorna el mismo state que le envié. Validar que sea el mismo. Viene en la query de request que envía okta
	if request.URL.Query().Get("state") != state {
		fmt.Fprintln(response, "Unexpected state") // Genera response con ese comment
		return
	}

	// Me retorna un code. Validar que sea no esté vacío. Viene en la query de request que envía okta
	if request.URL.Query().Get("code") == "" {
		fmt.Fprintln(response, "Unexpected code") // Genera response con ese comment
		return
	}

	// Pido toquen a Okta
	exchange := utils.ExchangeCode(request.URL.Query().Get("code"), request)
	if exchange.Error != "" {
		fmt.Println(exchange.Error)
		fmt.Println(exchange.ErrorDescription)
		return
	}

	// TODO: validate token
	//fmt.Println("Toquen obtenido", exchange.IdToken)

	// Guardo token en sesión
	session, err := sessionStore.Get(request, SESSION_KEY)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError) // Genero error response
	}

	session.Values["id_token"] = exchange.IdToken
	session.Values["access_token"] = exchange.AccessToken
	session.Save(request, response)
}

func RemoveSession(response http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, SESSION_KEY)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	delete(session.Values, "id_token")
	delete(session.Values, "access_token")
	session.Save(request, response)
}

func IsAuthenticated(request *http.Request) bool {
	session, err := sessionStore.Get(request, SESSION_KEY)

	if err != nil || session.Values["id_token"] == nil || session.Values["id_token"] == "" {
		return false
	}
	return true
}
