package memorystorage

import (
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	"sync"
)

type Storage struct {
	mu         sync.RWMutex
	Event      storage.Event
	EventStore map[string]*storage.Event
}

func New() *Storage {
	return &Storage{
		EventStore: make(map[string]*storage.Event),
	}
}

func (s *Storage) CreateEvent(event storage.Event) error {
	s.EventStore[event.ID] = &event

	return nil
}

func (s *Storage) UpdateEvent(event storage.Event) error {
	return nil
}

func (s *Storage) DeleteEvent(id string) error {
	return nil
}

func (s *Storage) GetEvents() map[string]*storage.Event {
	return s.EventStore
}

func (s *Storage) GetEvent(id string) error {
	return nil
}
