package stritten

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "Access-Control-Allow-Origin"},
		ExposedHeaders: []string{
			"Set-Cookie",
			"Authorization",
		},
		AllowCredentials: true,
		Debug:            true,
	})

	corsHandler := cors.Handler(handler)

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        corsHandler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
