package storage

import (
	"time"
)

type Event struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	DateTime    time.Time     `json:"dateTime"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
	UserID      string        `json:"userId"`
	NotifyIn    int           `json:"notifyIn"`
}

type Notification struct {
	ID           string
	Title        string
	DateTime     time.Time
	NotifiableID string
}
