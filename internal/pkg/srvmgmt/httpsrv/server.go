package httpsrv

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github/kunhou/gl_exercise/internal/pkg/srvmgmt"
)

const DefaultShutdownTimeout = 5 * time.Minute

var _ srvmgmt.Server = (*Server)(nil)

type Server struct {
	srv             *http.Server
	shutdownTimeout time.Duration
}

type Options func(*Server)

// NewServer creates and initializes a new Server with the provided Gin engine, configuration and optional settings.
// e: Gin Engine to handle HTTP requests.
// cfg: Configuration settings for the server.
// options: Optional functional parameters to modify the Server.
func NewServer(e *gin.Engine, cfg *Config, options ...Options) *Server {
	srv := Server{
		shutdownTimeout: DefaultShutdownTimeout,
		srv: &http.Server{
			Addr:    cfg.Addr,
			Handler: e,
		},
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv
}

// WithShutdownTimeout creates an option to set a custom timeout for graceful server shutdown.
// duration: Duration to wait for active connections to finish before shutting down.
func WithShutdownTimeout(duration time.Duration) Options {
	return func(server *Server) {
		server.shutdownTimeout = duration
	}
}

// Start launches the HTTP server. It returns an error if the server fails to start.
func (s *Server) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown gracefully stops the server. It waits for the configured timeout duration for active connections to close.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
