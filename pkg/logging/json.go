package logging

import (
	"errors"
	"io"
	"log/slog"
)

func UseJSONLogger(ws ...io.Writer) error {
	if len(ws) == 0 {
		return errors.New("UseJSONLogger must be called with at least one writter")
	}

	mw := io.MultiWriter(ws...)

	logger := slog.New(slog.NewJSONHandler(mw, nil))
	slog.SetDefault(logger)

	return nil
}
