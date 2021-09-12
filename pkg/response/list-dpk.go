package response

import "time"

type ListDpkResponse struct {
	Periode          time.Time `json:"periode" db:"periode"`
	NamaDebitur      string    `json:"nama_debitur" db:"nama_debitur"`
	Alamat           string    `json:"alamat" db:"alamat"`
	BakiDebit        float64   `json:"baki_debit" db:"baki_debit"`
	AvailableBalance string    `json:"available_balance" db:"available_balance"`
	NextPmtDate      string    `json:"next_pmt_date" db:"next_pmt_date"`
	FlagRestruk      string    `json:"flag_restruk" db:"flag_restruk"`
	LnType           string    `json:"ln_type" db:"ln_type"`
	NomorRekening    string    `json:"nomor_rekening" db:"nomor_rekening"`
	AccountNumber    string    `json:"account_number" db:"account_number"`
	Status           string    `json:"status" db:"status"`
	Billing          float64   `json:"billing" db:"billing"`
}
