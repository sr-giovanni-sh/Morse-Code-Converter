package server

import (
	"log"
	"net/http"
	"time"

	"Morse-Code-Converter/internal/handlers"
)

type Server struct {
	Log    *log.Logger
	Server *http.Server
}

// NewServer creates a new router and sets up the HTTP server configuration.
func NewServer(l *log.Logger) *Server {
	r := http.NewServeMux()

	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/upload", handlers.UploadHandler)

	hSer := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{Log: l, Server: hSer}
}
