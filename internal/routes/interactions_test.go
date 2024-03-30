package routes_test

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	"github.com/kytnacode/simple-discord-bot/internal/routes"
	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func TestInteractionsHandler_NoPublicKeyNoSignature(t *testing.T) {
	t.Setenv("BOT_PUBLIC_KEY", "")

	reqBody := discord.Dto{
		Type: 1, // Ping
	}

	rawBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Errorf("body couldn't be marshal: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "/interactions", bytes.NewBuffer(rawBody))
	if err != nil {
		t.Fatalf("request couldn't be created: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.InteractionsHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf(
			"server must return an error when no have public key, signature and/or timestamp: expected %v got: %v",
			http.StatusUnauthorized,
			rr.Code,
		)
	}
}

func TestInteractionsHandler_ValidRequest(t *testing.T) {
	pk, prk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("keys couldn't be generated: %v", err)
	}

	t.Setenv("BOT_PUBLIC_KEY", hex.EncodeToString(pk))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	reqBody := discord.Dto{
		Type: 1, // Ping
	}

	rawBody, err := json.Marshal(&reqBody)
	if err != nil {
		t.Errorf("body couldn't be marshal: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "/interactions", bytes.NewBuffer(rawBody))
	if err != nil {
		t.Errorf("request couldn't be created: %v", err)
	}

	timestamp := []byte("my timestamp")
	sg := ed25519.Sign(prk, slices.Concat(timestamp, rawBody))

	req.Header.Set("X-Signature-Ed25519", hex.EncodeToString(sg))
	req.Header.Set("X-Signature-Timestamp", string(timestamp))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.InteractionsHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("non 200 status code: expected %v got %v", http.StatusOK, rr.Code)
	}
}
