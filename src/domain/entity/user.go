package entity

import "time"

type User struct {
	ID        int
	UUID      string
	Username  string
	Password  string
	Limit     int
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (u *User) DaysLeft() int {
	return int(time.Until(u.ExpiresAt).Hours() / 24)
}
