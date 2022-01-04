package app

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	logger  *zap.Logger
	storage Storage
	config  *config.AppConfig
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
}

type Storage interface {
	CreateEvent(event storage.Event) (*storage.Event, error)
	UpdateEvent(event storage.Event) (storage.Event, error)
	DeleteEvent(id string) error
	GetEvents() []*storage.Event
	GetEvent(id string) *storage.Event
}

func New(logger *zap.Logger, storage Storage, config *config.AppConfig) *App {
	return &App{
		logger:  logger,
		storage: storage,
		config:  config,
	}
}

func (a *App) CreateEvent(ctx context.Context, data []byte) *storage.Event {
	event := storage.Event{}

	event.UserId = "change me"

	err := json.Unmarshal(data, &event)

	if err != nil {
		a.logger.Error("failed to unmarshal json")
	}

	created, err := a.storage.CreateEvent(event)

	if err != nil {
		a.logger.Error(err.Error())
	}

	ctx.Done()

	return created
}

func (a App) ListEvents(ctx context.Context) []*storage.Event {
	return a.storage.GetEvents()
}
