package types

import (
	"time"
)

type Event struct {
	ID      		string   	`json:"id"`
	Title    		string 		`json:"title"`
	Info    		string 		`json:"info"`
	CreatedAt 		time.Time	`json:"created"`
	UpdatedAt 		time.Time	`json:"updated"`
}

// BeforeInsert set CreatedAt and UpdatedAt.
func (e *Event) BeforeInsert() error {
	e.CreatedAt = time.Now().UTC().Truncate(time.Second)
	e.UpdatedAt = e.CreatedAt
	return nil
}

// BeforeUpdate set UpdatedAt.
func (e *Event) BeforeUpdate() error {
	e.UpdatedAt = time.Now().UTC().Truncate(time.Second)
	return nil
}
