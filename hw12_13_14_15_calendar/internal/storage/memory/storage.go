package memorystorage

import (
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

func (s *Storage) CreateEvent(event storage.Event) (*storage.Event, error) {
	event.ID = s.LastID
	s.LastID++
	s.EventStore[event.ID] = &event

	return &event, nil
}

func (s *Storage) UpdateEvent(event storage.Event) (*storage.Event, error) {
	// ineffectual assignment to existing (ineffassign) WTF?
	existing, ok := s.EventStore[event.ID] //nolint

	if !ok {
		return nil, errors.New("no events found for update")
	}

	existing = &event

	return existing, nil
}

func (s *Storage) DeleteEvent(id int) error {
	delete(s.EventStore, id)

	return nil
}

func (s *Storage) GetEvents() []*storage.Event {
	eventsSlice := make([]*storage.Event, 0, len(s.EventStore))

	for _, event := range s.EventStore {
		eventsSlice = append(eventsSlice, event)
	}

	return eventsSlice
}

func (s *Storage) GetEvent(id int) *storage.Event {
	return s.EventStore[id]
}
