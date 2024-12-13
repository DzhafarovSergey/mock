package config

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/lpernett/godotenv"
)

type ServiceConfig struct {
	ClientId  string
	SecretKey string
	CertPath  string
	KeyPath   string
}

func NewServiceConfig() *ServiceConfig {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Ошибка загрузки .env файла")
	}

	clientId, clientIdExist := os.LookupEnv("APP_CLIENT_ID")
	secretKey, secretKeyExist := os.LookupEnv("APP_SECRET_KEY")
	certPath, certPathExist := os.LookupEnv("APP_CERT_PATH")
	keyPath, keyPathExist := os.LookupEnv("APP_KEY_PATH")

	if !clientIdExist {
		clientId = "default_name"
	}
	if !secretKeyExist {
		secretKey = "default_name"
	}
	if !certPathExist {
		certPath = "default_name"
	}
	if !keyPathExist {
		keyPath = "default_name"
	}

	config := &ServiceConfig{
		ClientId:  clientId,
		SecretKey: secretKey,
		CertPath:  certPath,
		KeyPath:   keyPath,
	}

	return config
}
