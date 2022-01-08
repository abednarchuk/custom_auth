package models

import "time"

type User struct {
	ID         string
	Name       string
	UserName   string
	Email      string
	Password   string
	CreatedAt  time.Time
	ModifietAt time.Time
}
