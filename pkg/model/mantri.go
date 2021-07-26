package model

import "time"

type Mantri struct {
	ID          int       `json:"id" db:"id"`
	IDCabang    int       `json:"id_cabang" db:"id_cabang"`
	UsernamePN  string    `json:"username_pn" db:"username_pn"`
	NamaMantri  string    `json:"nama_mantri" db:"nama_mantri"`
	UnitKerja   string    `json:"unit_kerja" db:"unit_kerja"`
	KodeBranch  string    `json:"kode_branch" db:"kode_branch"`
	Jabatan     string    `json:"jabatan" db:"jabatan"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
