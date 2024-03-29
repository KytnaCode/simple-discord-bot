package routes

import (
	"fmt"
	"log/slog"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write([]byte("{ \"message\": \"Hello world\" }"))
	if err != nil {
		slog.Error("Error has occcurred on GET /: ", fmt.Errorf("Error writing response: %w", err))
		http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)

		return
	}
}
