package model

import "time"

type AksiMisi struct {
	ID          int       `json:"id" db:"id"`
	IDAksi      int       `json:"id_aksi" db:"id_aksi"`
	IDMisi      int       `json:"id_misi" db:"id_misi"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
