package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kytnacode/simple-discord-bot/internal/interactions"
)

var (
	// nolint:gochecknoglobals
	registerFlag = flag.Bool("register", false, "Register interactions in <root>/interactions")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if *registerFlag {
		err = register(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func register(ctx context.Context) error {
	const interactionsPath = "./interactions"

	registerFunc := interactions.CreateRegisterFunc(os.Getenv("APP_ID"), os.Getenv("BOT_TOKEN"))

	err := interactions.RegisterFolder(
		ctx,
		registerFunc,
		interactionsPath,
	)
	if err != nil {
		return fmt.Errorf("commands couldn't be registered: %w", err)
	}

	return nil
}
