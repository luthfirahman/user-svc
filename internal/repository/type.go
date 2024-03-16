package repository

import "time"

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
