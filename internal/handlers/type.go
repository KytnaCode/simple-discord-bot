package handlers

import (
	"net/http"

	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

type InteractionHandler func(w http.ResponseWriter, r *discord.Dto)
