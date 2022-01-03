package app

import (
	"context"
	"encoding/json"
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	logger  Logger
	storage Storage
	config  *config.AppConfig
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
}

type Storage interface {
	CreateEvent(event storage.Event) error
	UpdateEvent(event storage.Event) error
	DeleteEvent(id string) error
	GetEvents() map[string]*storage.Event
	GetEvent(id string) error
}

func New(logger Logger, storage Storage, config *config.AppConfig) *App {
	return &App{
		logger:  logger,
		storage: storage,
		config:  config,
	}
}

func (a *App) CreateEvent(ctx context.Context, data []byte) error {
	event := storage.Event{}

	err := json.Unmarshal(data, &event)

	if err != nil {
		a.logger.Error("failed to unmarshal json")
	}

	ctx.Done()

	return a.storage.CreateEvent(event)
}

func (a App) ListEvents(ctx context.Context) map[string]*storage.Event {
	return a.storage.GetEvents()
}
