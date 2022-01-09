package sqlstorage

import (
	"context"
	"database/sql"
	"log"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	sql *sql.DB
	dsn string
}

func New(cfg *config.Config) *Storage {
	return &Storage{
		sql: &sql.DB{},
		dsn: cfg.Db.Dsn,
	}
}

func (s *Storage) Connect(ctx context.Context) (err error) {
	s.sql, err = sql.Open("pgx", s.dsn)

	if err != nil {
		log.Println(err)
	}

	return s.sql.PingContext(ctx)
}

func (s *Storage) Close() error {
	return s.sql.Close()
}

func (s *Storage) CreateEvent(event storage.Event) (*storage.Event, error) {

	return nil, nil
}

func (s *Storage) UpdateEvent(event storage.Event) (*storage.Event, error) {
	return &event, nil
}

func (s *Storage) DeleteEvent(id int) error {
	return nil
}

func (s *Storage) GetEvents() []*storage.Event {
	return nil
}

func (s *Storage) GetEvent(id int) *storage.Event {
	return nil
}
