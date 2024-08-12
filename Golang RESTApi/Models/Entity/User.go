package entity

import "time"

type User struct {
	Id                        int
	Username, Email, Password string
	CreatedAt, UpdatedAt      time.Time
}

type UserProduct struct {
	Id              int
	Username, Email string
}
