package response

import "time"

type SisaLunasHutangResponse struct {
	Count           int        `json:"count" db:"count"`
	SisaLunasHutang float64    `json:"sisa_lunas_hutang" db:"sisa_lunas_hutang"`
	PnPengelola     string     `json:"pn_pengelola" db:"pn_pengelola"`
	TglRealisasi    *time.Time `json:"tgl_realisasi" db:"tgl_realisasi"`
	CifNo           string     `json:"cif_no" db:"cif_no"`
	Periode         *time.Time `json:"periode" db:"periode"`
	Branch          string     `json:"branch" db:"branch"`
	NomorRekening   string     `json:"nomor_rekening" db:"nomor_rekening"`
}
