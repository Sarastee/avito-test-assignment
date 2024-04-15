package main

import (
	"context"
	"flag"
	"log"

	"github.com/sarastee/avito-test-assignment/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()
}

// @title Banner Service
// @version 1.0.0
// @description Сервис для управления баннерами
// @schemes http

// @contact.name Ilya Lyakhov
// @contact.email ilja.sarasti@mail.ru

// @host localhost:8082
// @BasePath /

// @securityDefinitions.apikey AdminToken
// @name token
// @in header
// @description Admin access token

// @securityDefinitions.apikey UserToken
// @name token
// @in header
// @description User access token

func main() {
	ctx := context.Background()

	application, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatalf("init app failure: %s", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("failure while running the application: %s", err)
	}
}
