package types

import (
	"time"
)

type User struct {
	ID      		string   	`json:"id"`
	Name    		string 		`json:"name"`
	Email			string 		`json:"email"`
	Verified   		bool		`json:"verified"`
	Deleted			bool		`json:"deleted"`
	RegisteredAt 	time.Time	`json:"registered"`
	UpdatedAt 		time.Time	`json:"updated"`
}

// BeforeInsert set RegisteredAt and UpdatedAt.
func (u *User) BeforeInsert() error {
	u.RegisteredAt = time.Now().UTC().Truncate(time.Second)
	u.UpdatedAt = u.RegisteredAt
	return nil
}

// BeforeUpdate set UpdatedAt.
func (u *User) BeforeUpdate() error {
	u.UpdatedAt = time.Now().UTC().Truncate(time.Second)
	return nil
}
