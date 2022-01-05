package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	uuid "github.com/nu7hatch/gouuid"
)

type Storage struct {
	mu         *sync.RWMutex
	Event      *storage.Event
	EventStore map[string]*storage.Event
}

func New() *Storage {
	return &Storage{
		mu:         &sync.RWMutex{},
		EventStore: make(map[string]*storage.Event),
	}
}

func (s *Storage) CreateEvent(event storage.Event) (*storage.Event, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	event.ID = id.String()
	event.DateTime = time.Now()

	s.EventStore[event.ID] = &event

	return &event, nil
}

func (s *Storage) UpdateEvent(event storage.Event) (*storage.Event, error) {
	existing, ok := s.EventStore[event.ID] // ineffectual assignment to existing (ineffassign) WTF?

	if !ok {
		return nil, errors.New("no events found for update")
	}

	existing = &event

	return existing, nil
}

func (s *Storage) DeleteEvent(id string) error {
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

func (s *Storage) GetEvent(id string) *storage.Event {
	return s.EventStore[id]
}
