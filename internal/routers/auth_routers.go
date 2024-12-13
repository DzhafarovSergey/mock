package routers

import (
	"mock/internal/handlers"
	"mock/internal/services"
	"net/http"
)

func InitAuthRouter() *http.ServeMux {
	authService := services.NewAuthService()
	oathHandler := handlers.NewOAuthHandler(authService)

	mux := http.NewServeMux()
	mux.HandleFunc("/sso/oauth/token", oathHandler.HandleAccessToken)
	return mux
}
