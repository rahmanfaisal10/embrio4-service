package model

import "time"

type Unit struct {
	ID          int       `json:"id" db:"id"`
	IDCabang    int       `json:"id_cabang" db:"id_cabang"`
	Kode        string    `json:"kode" db:"kode"`
	Nama        string    `json:"nama" db:"nama"`
	Alamat      string    `json:"alamat" db:"alamat"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
