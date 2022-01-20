package memorystorage

import (
	"context"
	"errors"
	"sync"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu         *sync.RWMutex
	Event      *storage.Event
	EventStore map[int]*storage.Event
	LastID     int
}

func New() *Storage {
	return &Storage{
		mu:         &sync.RWMutex{},
		EventStore: make(map[int]*storage.Event),
		LastID:     1,
	}
}

func (s *Storage) CreateEvent(event storage.Event, ctx *context.Context) (*storage.Event, error) {
	event.ID = s.LastID
	s.LastID++
	s.mu.Lock()
	s.EventStore[event.ID] = &event
	s.mu.Unlock()

	return &event, nil
}

func (s *Storage) UpdateEvent(event storage.Event) (*storage.Event, error) {
	// ineffectual assignment to existing (ineffassign) WTF?
	s.mu.RLock()
	existing, ok := s.EventStore[event.ID] //nolint
	s.mu.RUnlock()

	if !ok {
		return nil, errors.New("no events found for update")
	}

	existing = &event

	return existing, nil
}

func (s *Storage) DeleteEvent(id int) error {
	s.mu.Lock()
	delete(s.EventStore, id)
	s.mu.Unlock()

	return nil
}

func (s *Storage) GetEvents() []*storage.Event {
	eventsSlice := make([]*storage.Event, 0, len(s.EventStore))

	s.mu.RLock()
	for _, event := range s.EventStore {
		eventsSlice = append(eventsSlice, event)
	}
	s.mu.RUnlock()

	return eventsSlice
}

func (s *Storage) GetEvent(id int) *storage.Event {
	return s.EventStore[id]
}
