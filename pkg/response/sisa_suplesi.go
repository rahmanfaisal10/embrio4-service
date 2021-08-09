package response

import "time"

type SisaSuplesiResponse struct {
	Count                   int       `json:"count" db:"count"`
	Periode                 time.Time `json:"periode" db:"periode"`
	PeriodeBulanSebelumnya  time.Time `json:"periode_bulan_sebelumnya" db:"periode_bulan_sebelumnya"`
	CIFNO                   string    `json:"CIFNO" db:"CIFNO"`
	NomorRekening           string    `json:"nomor_rekening" db:"nomor_rekening"`
	NomorRekeningSebelumnya string    `json:"nomor_rekening_sebelumnya" db:"nomor_rekening_sebelumnya"`
	SisaSuplesi             float64   `json:"sisa_suplesi" db:"sisa_suplesi"`
}
