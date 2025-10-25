package domain

import "time"

type Notification struct {
	ID        string
	Type      string
	Recipient string
	Subject   string
	Body      string
	SentAt    time.Time
	CreatedAt time.Time
}
