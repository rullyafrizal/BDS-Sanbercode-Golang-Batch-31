package models

import "time"

type Student struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
