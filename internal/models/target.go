package models

import "time"

type Target struct {
	ID         int       `json:"id"`
	MissionID  uint      `json:"mission_id,omitempty"`
	Name       string    `json:"name,omitempty" `
	Country    string    `json:"country,omitempty" `
	Notes      string    `json:"notes,omitempty"`
	IsComplete bool      `json:"is_complete,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
}
