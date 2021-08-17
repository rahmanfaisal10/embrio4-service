package model

import "time"

type Simpanan struct {
	ID                     int64      `json:"id" db:"id"`
	Periode                *time.Time `json:"periode" db:"periode"`
	UkerCode               string     `json:"uker_code" db:"uker_code"`
	CurrCode               string     `json:"curr_code" db:"curr_code"`
	CUrrDesc               string     `json:"curr_desc" db:"curr_desc"`
	AccountNumber          string     `json:"account_number" db:"account_number"`
	CifNo                  string     `json:"cif_no" db:"cif_no"`
	ShortName              string     `json:"short_name" db:"short_name"`
	OfficerName            string     `json:"officer_name" db:"officer_name"`
	Dlt                    *time.Time `json:"dlt" db:"dlt"`
	OpenDT                 *time.Time `json:"open_dt" db:"open_dt"`
	Balance                float64    `json:"balance" db:"balance"`
	AvailableBalance       float64    `json:"available_balance" db:"available_balance"`
	IntCredit              float64    `json:"int_credit" db:"int_credit"`
	AccruedInt             float64    `json:"accrued_int" db:"accrued_int"`
	AverageBalance         float64    `json:"average_balance" db:"average_balance"`
	ProdCode               string     `json:"prod_code" db:"prod_code"`
	PnPengelola            int64      `json:"pn_pengelola" db:"pn_pengelola"`
	NamaPengelola          string     `json:"nama_pengelola" db:"nama_pengelola"`
	KecamatanTempatTinggal string     `json:"kecamatan_tempat_tinggal" db:"kecamatan_tempat_tinggal"`
	KelurahanTempatTinggal string     `json:"kelurahan_tempat_tinggal" db:"kelurahan_tempat_tinggal"`
	KodePosTempatTinggal   string     `json:"kode_pos_tempat_tinggal" db:"kode_pos_tempat_tinggal"`
	KecamatanTempatUsaha   string     `json:"kecamatan_tempat_usaha" db:"kecamatan_tempat_usaha"`
	KelurahanTempatUsaha   string     `json:"kelurahan_tempat_usaha" db:"kelurahan_tempat_usaha"`
	KodePosTempatUsaha     string     `json:"kode_pos_tempat_usaha" db:"kode_pos_tempat_usaha"`
}
