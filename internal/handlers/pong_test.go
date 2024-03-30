package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kytnacode/simple-discord-bot/internal/handlers"
	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func pongHandler(w http.ResponseWriter, _ *http.Request) {
	handlers.Pong(w, nil)
}

func TestPong_ShouldReturnANonEmptyBody(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "/interactions", nil)
	if err != nil {
		t.Fatalf("request couldn't be created: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pongHandler)

	handler.ServeHTTP(rr, req)

	if len(rr.Body.Bytes()) == 0 {
		t.Error("pong must return a non-empty body")
	}
}

func TestPong_ShouldReturnABodyWithTypeOne(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", "/interactions", nil)
	if err != nil {
		t.Fatalf("request couldn't be created: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pongHandler)

	handler.ServeHTTP(rr, req)

	var dto discord.Dto

	err = json.Unmarshal(rr.Body.Bytes(), &dto)
	if err != nil {
		t.Fatalf("response couldn't be unmarshal: %v", err)
	}

	if dto.Type != 1 {
		t.Errorf("dto.Type should be 1: expected %v got %v", 1, dto.Type)
	}
}
