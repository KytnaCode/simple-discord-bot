package discord_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func TestCommandBuilder_ShouldNotReturnAnError(t *testing.T) {
	t.Parallel()

	expected := "{ \"my\": \"response\" }"

	cb := new(discord.CommandBuilder)

	cb.SetAppID("my app id")
	cb.SetToken("my token")
	cb.SetJSON([]byte("{ \"my\": \"json\" }"))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_, err := w.Write([]byte(expected))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	cb.SetURL(server.URL)

	b, err := cb.RegisterJSON(context.Background())
	if err != nil {
		t.Errorf("RegisterJSON failed: %v", err)
	}

	if string(b) != expected {
		t.Errorf("Expected %v got %v", expected, string(b))
	}
}
