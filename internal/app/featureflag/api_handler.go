package featureflag

import (
	"fmt"
	"net/http"
	"snapp-featureflag/internal/app/featureflag/feature"
)

type ApiHandler struct {
	commandHandler feature.CommandHandler
	queryHandler   feature.QueryHandler
}

func NewApiHandler(commandHandler feature.CommandHandler, queryHandler feature.QueryHandler) *ApiHandler {
	return &ApiHandler{commandHandler, queryHandler}
}

func (a *ApiHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
}
