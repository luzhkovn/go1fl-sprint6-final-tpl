package server

import (
	"log"
	"net/http"
	"time"

	"github.com/luzhkovn/go1fl-sprint6-final-tpl/internal/handlers"
)

type Server struct {
	logger *log.Logger
	server http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	return &Server{
		logger: logger,
		server: http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

}

func (s *Server) Start() error {
	err := s.server.ListenAndServe()
	return err
}
