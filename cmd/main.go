package main

import (
	"flag"
	"fmt"
	"mock/internal/middleware"
	"mock/internal/routers"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	var ip string
	var port int

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.DateTime,
		FullTimestamp:   true,
		ForceColors:     true,
	})

	flag.StringVar(&ip, "ip", "127.0.0.1", "IP адрес")
	flag.IntVar(&port, "port", 8000, "Порт")
	flag.Parse()

	var authRouter = routers.InitAuthRouter()
	var authLoggerRouter = middleware.LoggingMiddleware(authRouter)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/sso/", authLoggerRouter)

	var addr = fmt.Sprintf("%s:%d", ip, port)

	log.Printf("Запуск http сервера на адресе %s", addr)
	if err := http.ListenAndServe(addr, mainRouter); err != nil {
		log.Fatalf("Не удалось запустить серверс %s", err.Error())
	}

}
