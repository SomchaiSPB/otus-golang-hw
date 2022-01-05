package internalhttp

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	app    Application
	config *config.AppConfig
}

type Application interface {
	CreateEvent(ctx context.Context, data []byte) *storage.Event
	ListEvents(ctx context.Context) []*storage.Event
}

func NewServer(logger *zap.Logger, app Application, config *config.AppConfig) *Server {
	return &Server{
		logger: logger,
		app:    app,
		config: config,
	}
}

func (s *Server) Start(ctx context.Context) error {
	address := net.JoinHostPort(s.config.Host, s.config.Port)

	server := &http.Server{Addr: address, Handler: s.service()}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(ctx)

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				s.logger.Error("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			s.logger.Error(err.Error())
		}
		cancel()
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		s.logger.Error(err.Error())
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Server) service() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(s.loggingMiddleware)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("failed to read post body"))
			return
		}

		response := s.app.CreateEvent(context.Background(), data)
		if err != nil {
			w.Write([]byte("failed to create event"))
			return
		}

		responseData, err := json.Marshal(response)
		if err != nil {
			s.logger.Error(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(responseData)
	})

	r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
		var response []storage.Event
		events := s.app.ListEvents(context.Background())

		for _, event := range events {
			response = append(response, *event)
		}

		responseData, err := json.Marshal(response)
		if err != nil {
			s.logger.Error(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(responseData)
	})

	return r
}
