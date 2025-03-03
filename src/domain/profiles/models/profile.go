package models

import "time"

type Profile struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	Age         int64
	Job         string
	State       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StateEntity string
}
