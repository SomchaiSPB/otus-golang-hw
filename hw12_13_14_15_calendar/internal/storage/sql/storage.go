package sqlstorage

import (
	"context"
	"database/sql"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	sql sql.DB
}

func New(*config.Config) *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) CreateEvent(event storage.Event) (*storage.Event, error) {
	return nil, nil
}

func (s *Storage) UpdateEvent(event storage.Event) (*storage.Event, error) {
	return &event, nil
}

func (s *Storage) DeleteEvent(id string) error {
	return nil
}

func (s *Storage) GetEvents() []*storage.Event {
	return nil
}

func (s *Storage) GetEvent(id string) *storage.Event {
	return nil
}
