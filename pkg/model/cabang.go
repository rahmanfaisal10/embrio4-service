package model

import "time"

type Cabang struct {
	ID          int       `json:"id" db:"id"`
	Cabang      string    `json:"cabang" db:"cabang"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
