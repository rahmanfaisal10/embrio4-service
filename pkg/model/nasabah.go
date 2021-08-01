package model

import "time"

type Nasabah struct {
	ID               int       `json:"id" db:"id"`
	Nama             string    `json:"nama" db:"nama"`
	NoKtp            string    `json:"no_ktp" db:"no_ktp"`
	NomorRekening    string    `json:"nomor_rekening" db:"nomor_rekening"`
	CifNo            string    `json:"cif_no" db:"cif_no"`
	Saldo            float64   `json:"saldo" db:"saldo"`
	NoTelepon        string    `json:"no_telepon" db:"no_telepon"`
	NamaIbuKandung   string    `json:"nama_ibu_kandung" db:"nama_ibu_kandung"`
	KodePosIdentitas string    `json:"kode_pos_identitas" db:"kode_pos_identitas"`
	AlamatIdentitas  string    `json:"alamat_identitas" db:"alamat_identitas"`
	KodePosKantor    string    `json:"kode_pos_kantor" db:"kode_pos_kantor"`
	AlamatKantor     string    `json:"alamat_kantor" db:"alamat_kantor"`
	Description      string    `json:"description" db:"description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
