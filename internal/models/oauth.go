package models

type OAuthRequest struct {
	ClientId  string `json:"clientId"`
	SecretKey string `json:"secretKey"`
}

type OAuthResponse struct {
	Token string `json:"token"`
}
