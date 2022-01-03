package storage

import "time"

type Event struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	DateTime    time.Time     `json:"dateTime"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
	UserId      string        `json:"userId"`
	NotifyIn    int           `json:"notifyIn"`
}

type Notification struct {
	Id           string
	Title        string
	DateTime     time.Time
	NotifiableId string
}
