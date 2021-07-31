package model

import "time"

type DetailTunggakan struct {
	ID               int       `json:"id" db:"id"`
	IDDetailCicilan  int       `json:"id_detail_cicilan" db:"id_detail_cicilan"`
	IDJenisTunggakan int       `json:"id_jenis_tunggakan" db:"id_jenis_tunggakan"`
	Nominal          float64   `json:"nominal" db:"nominal"`
	Description      string    `json:"description" db:"description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
