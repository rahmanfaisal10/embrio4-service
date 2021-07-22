package model

import "time"

type Users struct {
	ID          int       `json:"id" db:"id"`
	UsernamePN  string    `json:"username_pn" db:"username_pn"`
	Password    string    `json:"password" db:"password"`
	Nama        string    `json:"nama" db:"nama"`
	UnitKerja   string    `json:"unit_kerja" db:"unit_kerja"`
	KodeBranch  string    `json:"kode_branch" db:"kode_branch"`
	Jabatan     string    `json:"jabatan" db:"jabatan"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
	LastLogin   time.Time `json:"last_login" db:"last_login"`
	Groups      string    `json:"groups" db:"groups"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
}
