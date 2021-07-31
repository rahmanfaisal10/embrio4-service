package model

import "time"

type Nasabah struct {
	ID                     int       `json:"id" db:"id"`
	Nama                   string    `json:"nama" db:"nama"`
	NoKtp                  string    `json:"no_ktp" db:"no_ktp"`
	NomorRekening          string    `json:"nomor_rekening" db:"nomor_rekening"`
	Saldo                  float64   `json:"saldo" db:"saldo"`
	NoTelepon              string    `json:"no_telepon" db:"no_telepon"`
	NamaIbuKandung         string    `json:"nama_ibu_kandung" db:"nama_ibu_kandung"`
	KodePosTempatTinggal   string    `json:"kode_pos_tempat_tinggal" db:"kode_pos_tempat_tinggal"`
	KelurahanTempatTinggal string    `json:"kelurahan_tempat_tinggal" db:"kelurahan_tempat_tinggal"`
	KecamatanTempatTinggal string    `json:"kecamatan_tempat_tinggal" db:"kecamatan_tempat_tinggal"`
	KodePosTempatUsaha     string    `json:"kode_pos_tempat_usaha" db:"kode_pos_tempat_usaha"`
	KelurahanTempatUsaha   string    `json:"kelurahan_tempat_usaha" db:"kelurahan_tempat_usaha"`
	KecamatanTempatUsaha   string    `json:"kecamatan_tempat_usaha" db:"kecamatan_tempat_usaha"`
	Description            string    `json:"description" db:"description"`
	CreatedAt              time.Time `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time `json:"updated_at" db:"updated_at"`
}
