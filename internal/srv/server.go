package srv

import (
	"log/slog"
	"net/http"
)

func NewServer(mux *http.ServeMux) *http.Server {
	return &http.Server{
		ReadTimeout:                  readTimeout,
		WriteTimeout:                 writeTimeout,
		IdleTimeout:                  idleTimeout,
		ReadHeaderTimeout:            readHeaderTimeout,
		DisableGeneralOptionsHandler: disableGeneralOptionsHandler,
		Addr:                         addr,
		Handler:                      mux,
		MaxHeaderBytes:               maxHeaderBytes,
		ErrorLog:                     slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
		TLSConfig:                    nil,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ConnContext:                  nil,
		BaseContext:                  nil,
	}
}
