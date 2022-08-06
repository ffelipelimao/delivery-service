package domain

import "time"

type Register struct {
	ID          string
	From        string
	Destination string
	CreatedAt   time.Time
	Objects     []Object
}
