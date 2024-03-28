package main

import (
	"log/slog"
	"os"

	"github.com/kytnacode/simple-discord-bot/pkg/logging"
)

func main() {
	// Open a the log file writter.
	f, err := os.Create("log.json") // Log file's name.
	if err != nil {
		slog.Default().Error("Error has occurred: ", err)

		return
	}
	defer f.Close()

	// Make default slog json to write to `os.Stdout` and a log file in json format.
	err = logging.UseJSONLogger(os.Stdout, f)
	if err != nil {
		slog.Default().Error("Error has occurred: ", err)

		return
	}

	// Get new json logger.
	logger := slog.Default()

	logger.Info("Starting App!")
}
