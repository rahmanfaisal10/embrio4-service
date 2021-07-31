package model

import "time"

type DetailCicilan struct {
	ID               int       `json:"id" db:"id"`
	IDCicilan        int       `json:"id_cicilan" db:"id_cicilan"`
	IDKolektebilitas int       `json:"id_kolektebilitas" db:"id_kolektebilitas"`
	IDAksiMisi       int       `json:"id_aksi_misi" db:"id_aksi_misi"`
	TanggalBayar     time.Time `json:"tanggal_bayar" db:"tanggal_bayar"`
	TglJanjiBayar    time.Time `json:"tgl_janji_bayar" db:"tgl_janji_bayar"`
	Keterangan       string    `json:"keterangan" db:"keterangan"`
	Description      string    `json:"description" db:"description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
