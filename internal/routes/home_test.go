package routes_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kytnacode/simple-discord-bot/internal/routes"
)

func TestHome_ShouldReturnAnOkResponse(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.Home)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestHome_ShouldReturnAValidJSON(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.Home)

	handler.ServeHTTP(rr, req)

	if !json.Valid(rr.Body.Bytes()) {
		t.Errorf("home must return a valid json: got %v", rr.Body.String())
	}
}

func TestHome_ShouldNotMatchNonExactRootPath(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"/this-route-should-not-be-matched-by-home",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.Home)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf(
			"home handler must not match non exact path: path %v status %v expected status %v",
			req.URL,
			rr.Code,
			http.StatusNotFound,
		)
	}
}
