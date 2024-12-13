package routers

import (
	"mock/internal/handlers"
	"mock/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

func InitBaseRouter() *mux.Router {
	authService := services.NewAuthService()
	oauthHandler := handlers.NewOAuthHandler(authService)

	router := mux.NewRouter()

	router.HandleFunc("/sso/oauth/token", oauthHandler.HandleAccessToken).Methods(http.MethodPost)

	return router
}
