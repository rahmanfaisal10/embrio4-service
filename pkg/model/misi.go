package model

import "time"

type Misi struct {
	ID          int       `json:"id" db:"id"`
	IDMantri    int       `json:"id_mantri" db:"id_mantri"`
	Kode        string    `json:"kode" db:"kode"`
	Nama        string    `json:"nama" db:"nama"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
