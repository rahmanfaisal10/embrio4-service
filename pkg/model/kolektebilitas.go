package model

import "time"

type Kolektebilitas struct {
	ID          int       `json:"id" db:"id"`
	Kode        string    `json:"kode" db:"kode"`
	Nama        string    `json:"nama" db:"nama"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
