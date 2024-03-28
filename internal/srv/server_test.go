package srv_test

import (
	"net/http"
	"testing"

	"github.com/kytnacode/simple-discord-bot/internal/srv"
)

func TestNewServer_ShouldReturnANonNilServer(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()
	server := srv.NewServer(mux)

	if server == nil {
		t.Fatal("Server must not be nil")
	}
}

func TestNewServer_ShouldReturnAServerWithTheCorrectHandler(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()
	server := srv.NewServer(mux)

	if server.Handler != mux {
		t.Fatal("Server handler must be the argument's mux")
	}
}
