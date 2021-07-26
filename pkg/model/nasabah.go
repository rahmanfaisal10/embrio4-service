package model

import "time"

type Nasabah struct {
	ID                     int       `json:"id" db:"id"`
	IDMantri               int       `json:"id_mantri" db:"id_mantri"`
	NamaNasabah            string    `json:"nama_nasabah" db:"nama_nasabah"`
	NomorRekening          string    `json:"nomor_rekening" db:"nomor_rekening"`
	CifNo                  string    `json:"cif_no" db:"cif_no"`
	NoKtp                  string    `json:"no_ktp" db:"no_ktp"`
	KecamatanTempatTinggal string    `json:"kecamatan_tempat_tinggal" db:"kecamatan_tempat_tinggal"`
	KelurahanTempatTinggal string    `json:"kelurahan_tempat_tinggal" db:"kelurahan_tempat_tinggal"`
	KodePosTempatTinggal   string    `json:"kode_pos_tempat_tinggal" db:"kode_pos_tempat_tinggal"`
	KecamatanTempatUsaha   string    `json:"kecamatan_tempat_usaha" db:"kecamatan_tempat_usaha"`
	KelurahanTempatUsaha   string    `json:"kelurahan_tempat_usaha" db:"kelurahan_tempat_usaha"`
	KodePosTempatUsaha     string    `json:"kode_pos_tempat_usaha" db:"kode_pos_tempat_usaha"`
	Description            string    `json:"description" db:"description"`
	CreatedAt              time.Time `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time `json:"updated_at" db:"updated_at"`
}
