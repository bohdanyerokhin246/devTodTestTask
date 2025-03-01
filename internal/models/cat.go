package models

import "time"

type Cat struct {
	ID         int       `json:"id"`
	Name       string    `json:"name,omitempty"`
	Experience int       `json:"experience,omitempty"`
	Breed      string    `json:"breed,omitempty"`
	Salary     float64   `json:"salary,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
}
