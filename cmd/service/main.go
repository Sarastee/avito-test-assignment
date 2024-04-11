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
