package model

import "time"

type Dashboard struct {
	ID                 int       `json:"id" db:"db"`
	Periode            time.Time `json:"periode" db:"periode"`
	TargetOs           float64   `json:"target_os" db:"target_os"`
	PosisiOs           float64   `json:"posisi_os" db:"posisi_os"`
	PercenPencapaianOs float64   `json:"percen_pencapaian_os" db:"percen_pencapaian_os"`
}
