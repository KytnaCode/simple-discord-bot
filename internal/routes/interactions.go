package routes

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/kytnacode/simple-discord-bot/internal/bot"
	"github.com/kytnacode/simple-discord-bot/internal/handlers"
	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func handleHexError(err error, w http.ResponseWriter) bool {
	if err != nil {
		slog.Error("error on InteractionsHandler: Decode hex", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return false
	}

	return true
}

func InteractionsHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error on InteractionsHandler: body read: ", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	publicKeyHex := bot.GetPublicKey()
	signatureHex := r.Header.Get("X-Signature-Ed25519")
	timestamp := r.Header.Get("X-Signature-Timestamp")

	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if ok := handleHexError(err, w); !ok {
		return
	}

	signatureBytes, err := hex.DecodeString(signatureHex)
	if ok := handleHexError(err, w); !ok {
		return
	}

	rv := new(discord.RequestVerifier)
	rv.SetPublicKey(publicKeyBytes)
	rv.SetSignature(signatureBytes)
	rv.SetTimestamp([]byte(timestamp))
	rv.SetBodyContent(b)

	isVerified := rv.Verify()

	if !isVerified {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return
	}

	var req discord.Dto

	err = json.Unmarshal(b, &req)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if req.Type == 1 {
		handlers.Pong(w, &req)
	}

	if req.Type == 2 {
		var interactionReq discord.InteractionRequest

		err = json.Unmarshal(b, &interactionReq)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		handlers.CommandHandler(w, interactionReq)
	}
}
