package main

import (
	"context"
	"log"

	"github.com/drizzleent/auth/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("Failed to init app: %v", err.Error())
	}

	err = app.Run()

	if err != nil {
		log.Fatalf("Failed to run app: %v", err.Error())
	}
}
