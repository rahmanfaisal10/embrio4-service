package model

import "time"

type Upload struct {
	ID                         int        `json:"id,omitempty"`
	Periode                    *time.Time `json:"periode" db:"periode"`
	Branch                     string     `json:"branch" db:"branch"`
	Currency                   string     `json:"currency" db:"currency"`
	NamaAO                     string     `json:"nama_ao" db:"nama_ao"`
	LNType                     string     `json:"ln_type" db:"ln_type"`
	NomorRekening              string     `json:"nomor_rekening" db:"nomor_rekening"`
	NamaDebitur                string     `json:"nama_debitur" db:"nama_debitur"`
	Plafond                    float64    `json:"plafond" db:"plafond"`
	NextPmtDate                *time.Time `json:"next_pmt_date" db:"next_pmt_date"`
	NextIntPmtDate             *time.Time `json:"next_int_pmt_date" db:"next_int_pmt_date"`
	Rate                       float32    `json:"rate" db:"rate"`
	TglMenunggak               *time.Time `json:"tgl_menunggak" db:"tgl_menunggak"`
	TglRealisasi               *time.Time `json:"tgl_realisasi" db:"tgl_realisasi"`
	TglJatuhTempo              *time.Time `json:"tgl_jatuh_tempo" db:"tgl_jatuh_tempo"`
	JangkaWaktu                string     `json:"jangka_waktu" db:"jangka_waktu"`
	FlagRestruk                string     `json:"flag_restruk" db:"flag_restruk"`
	CIFNO                      string     `json:"cif_no" db:"cif_no"`
	KolektibilitasLancar       float64    `json:"kolektibilitas_lancar" db:"kolektibilitas_lancar"`
	KolektibilitasDPK          float64    `json:"kolektibilitas_dpk" db:"kolektibilitas_dpk"`
	KolektibilitasKurangLancar float64    `json:"kolektibilitas_kurang_lancar" db:"kolektibilitas_kurang_lancar"`
	KolektibilitasDiragukan    float64    `json:"kOlektibilitas_diragukan" db:"kolektibilitas_diragukan"`
	KolektibilitasMacet        float64    `json:"kolektibilitas_macet" db:"kolektibilitas_macet"`
	TunggakanPokok             float64    `json:"tunggakan_pokok" db:"tunggakan_pokok"`
	TunggakanBunga             float64    `json:"tunggakan_bunga" db:"tunggakan_bunga"`
	TunggakanPinalty           float64    `json:"tunggakan_pinalty" db:"tunggakan_pinalty"`
	PNPengelola                string     `json:"pn_pengelola" db:"pn_pengelola"`
	NamaPengelola              string     `json:"nama_pengelola" db:"nama_pengelola"`
	Code                       string     `json:"code" db:"code"`
	Description                string     `json:"description" db:"description"`
	KolADK                     int64      `json:"kol_adk" db:"kol_adk"`
	AvgOsHarian                string     `json:"avg_os_harian" db:"avg_os_harian"`
	KecamatanTempatTinggal     string     `json:"kecamatan_tempat_tinggal" db:"kecamatan_tempat_tinggal"`
	KelurahanTempatTinggal     string     `json:"kelurahan_tempat_tinggal" db:"kelurahan_tempat_tinggal"`
	KodePosTempatTinggal       string     `json:"kode_pos_tempat_tinggal" db:"kode_pos_tempat_tinggal"`
	KecamatanTempatUsaha       string     `json:"kecamatan_tempat_usaha" db:"kecamatan_tempat_usaha"`
	KelurahanTempatUsaha       string     `json:"kelurahan_tempat_usaha" db:"kelurahan_tempat_usaha"`
	KodePosTempatUsaha         string     `json:"kode_pos_tempat_usaha" db:"kode_pos_tempat_usaha"`
}
