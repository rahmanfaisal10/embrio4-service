package model

import "time"

type Pinjaman struct {
	ID               int       `json:"id" db:"id"`
	IDMantri         int       `json:"id_mantri" db:"id_mantri"`
	IDJenisPinjaman  int       `json:"id_jenis_pinjaman" db:"id_jenis_pinjaman"`
	IDNasabah        int       `json:"id_nasabah" db:"id_nasabah"`
	Kode             string    `json:"kode" db:"kode"`
	Jumlah           float64   `json:"jumlah" db:"jumlah"`
	JangkaWaktu      int       `json:"jangka_waktu" db:"jangka_waktu"`
	TanggalPencairan time.Time `json:"tanggal_pencairan" db:"tanggal_pencairan"`
	Description      string    `json:"description" db:"description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
