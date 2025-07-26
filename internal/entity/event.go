package entity

import "time"

type Event struct {
	Id          int
	Name        string
	Description string
	Place       string
	Date        time.Time
	Count       int
	Fundraising int
	UserId      int
}
