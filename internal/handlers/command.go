package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/kytnacode/simple-discord-bot/internal/commands"
	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func CommandHandler(w http.ResponseWriter, req discord.InteractionRequest) {
	var reply discord.InteractionReply

	slog.Info("interaction", "name", req.Data.Name)

	if req.Data.Name == "ping" {
		reply = commands.PingHandler()
	}

	if req.Data.Name != "ping" {
		slog.Warn("Command registered but not handled: %v", "command", req.Data.Name)
		reply = discord.InteractionReply{
			Type: 1, // if command is unknown, return a pong response.
			Data: nil,
		}
	}

	b, err := json.Marshal(reply)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	slog.Info(string(b))
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
