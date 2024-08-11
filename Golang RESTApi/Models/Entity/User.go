package entity

import "time"

type User struct {
	Id                        int
	Username, Email, Password string
	CreatedAt, UpdatedAt      time.Time
}
