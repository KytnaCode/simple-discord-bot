package logging_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/kytnacode/simple-discord-bot/pkg/logging"
)

func getJSONLogger(t *testing.T, ws ...io.Writer) (*slog.Logger, error) {
	t.Helper()

	err := logging.UseJSONLogger(ws...)
	if err != nil {
		return nil, fmt.Errorf("logging.UseJSONLogger return an error: %w", err)
	}

	return slog.Default(), nil
}

func checkAndGetJSONLogger(t *testing.T, ws ...io.Writer) *slog.Logger {
	t.Helper()

	logger, err := getJSONLogger(t, ws...)
	if err != nil {
		t.Fatalf("UseJSONLogger returned an error: %v\n", err)
	}

	return logger
}

func TestUseJSONLogger_ShouldNotReturnAnError(t *testing.T) {
	t.Parallel()

	checkAndGetJSONLogger(t, os.Stdout)
}

func TestUseJSONLogger_ShouldReturnAnErrorIfWrittersListIsEmpty(t *testing.T) {
	t.Parallel()

	_, err := getJSONLogger(t)
	if err == nil {
		t.Fatal("UseJSONLogger should return an error if any writters aren't passed")
	}
}

func TestUseJSONLogger_ShouldChangeTheDefaultLogger(t *testing.T) {
	t.Parallel()

	oldLogger := slog.Default()
	jsonLogger := checkAndGetJSONLogger(t, os.Stdout)

	if oldLogger == jsonLogger {
		t.Fatal("slog.Default() logger should be a new JSON logger")
	}
}

func TestUseJSONLogger_ShouldUseTheFirstWritter(t *testing.T) {
	t.Parallel()

	oldValue := ""

	w := bytes.NewBufferString(oldValue)

	logger := checkAndGetJSONLogger(t, w)
	logger.Info("Hello World")

	newValue := w.String()

	if oldValue == newValue {
		t.Fatal("The writter has not the expected data\n")
	}
}

func TestUseJSONLogger_ShouldUseAllWritters(t *testing.T) {
	t.Parallel()

	var w1, w2, w3, w4 bytes.Buffer

	n := 4

	logger := checkAndGetJSONLogger(t, &w1, &w2, &w3, &w4)
	logger.Info("Hello World!")

	ws := []*bytes.Buffer{&w1, &w2, &w3, &w4}

	for i := range n {
		if ws[i].String() == "" {
			t.Fatalf("Writter number `%v` was not used", i)
		}
	}
}

func TestUseJSONLogger_shouldWriteAValidJSON(t *testing.T) {
	t.Parallel()

	var w bytes.Buffer

	logger := checkAndGetJSONLogger(t, &w)
	logger.Info("Hello World!")

	if !json.Valid(w.Bytes()) {
		t.Fatalf("logger should write valid json output")
	}
}
