package models

import "time"

type NewItem struct {
	User_id     string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type Item struct {
	Id          string    `json:"id"`
	User_id     string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
