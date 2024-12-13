package handlers

import (
	"encoding/json"
	"mock/internal/models"
	"mock/internal/services"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type OAuthHandler struct {
	authService *services.AuthService
}

func NewOAuthHandler(authService *services.AuthService) *OAuthHandler {
	return &OAuthHandler{authService: authService}
}

func (h *OAuthHandler) HandleAccessToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request models.OAuthRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println("Ошибка декодирования тела запроса", err)
	}

	token, err := h.authService.ValidateCredentials(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Println("Неудачная попытка авторизации для ClientId", request.ClientId)
		return
	}

	w.Header().Set("Content-Type", "applecation/json")
	w.Write([]byte(token))

}
