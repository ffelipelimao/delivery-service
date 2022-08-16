package domain

import "time"

type Register struct {
	ID          string    `json:"id" gorm:"type:uuid;primary_key"`
	From        string    `json:"from"`
	Destination string    `json:"destination"`
	CreatedAt   time.Time `json:"created_at"`
	Objects     []*Object `json:"objects"`
}

func (Register) TableName() string {
	return "register"
}
