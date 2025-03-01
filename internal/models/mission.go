package models

import "time"

type Mission struct {
	ID         uint      `json:"id"`
	CatID      uint      `json:"cat_id,omitempty"`
	IsComplete bool      `json:"is_complete,omitempty"`
	Targets    []Target  `json:"targets,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
}
