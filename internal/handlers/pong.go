package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func Pong(w http.ResponseWriter, _ *discord.Dto) {
	res := discord.Dto{Type: 1}

	resMarshal, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(resMarshal)
	if err != nil {
		slog.Error("error in pong handler: cannot write response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
