package storage

import (
	"time"
)

type Event struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
	UserID      int           `json:"userId"`
	NotifyIn    int           `json:"notifyIn"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type Notification struct {
	ID           string
	Title        string
	DateTime     time.Time
	NotifiableID string
}
