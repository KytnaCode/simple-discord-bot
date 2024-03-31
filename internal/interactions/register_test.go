package interactions_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kytnacode/simple-discord-bot/internal/interactions"
	"github.com/kytnacode/simple-discord-bot/pkg/discord"
)

func TestRegisterFolder_ShouldToNotReturnAnError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_, err := w.Write([]byte("{ \"my\": \"response\" }"))
		if err != nil {
			t.Error(err)
		}
	}))

	register := func(ctx context.Context, b []byte) ([]byte, error) {
		cb := new(discord.CommandBuilder)

		cb.SetToken("MY_TOKEN")
		cb.SetAppID("MY_APP_ID")
		cb.SetJSON(b)
		cb.SetURL(server.URL)

		return cb.RegisterJSON(ctx)
	}

	tmp := t.TempDir()

	f, err := os.CreateTemp(tmp, "*")
	if err != nil {
		t.Error(err)
	}

	fmt.Fprintln(f, "{ \"my\": \"file\" }")

	err = interactions.RegisterFolder(context.Background(), register, tmp)
	if err != nil {
		t.Error(err)
	}
}
