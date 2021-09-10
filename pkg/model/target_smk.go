package model

import "time"

type TargetSmk struct {
	ID               int64      `json:"id" db:"id"`
	NamaPN           string     `json:"nama_pn" db:"nama_pn"`
	PN               string     `json:"pn" db:"pn"`
	SegmentMantri    string     `json:"segment_mantri" db:"segment_mantri"`
	Tanggal          *time.Time `json:"tanggal" db:"tanggal"`
	TotalOS          float64    `json:"total_os" db:"total_os"`
	OsKupedes        float64    `json:"os_kupedes" db:"os_kupedes"`
	OsBriguna        int64      `json:"os_briguna" db:"os_briguna"`
	OsKur            float64    `json:"os_kur" db:"os_kur"`
	AverageSimpedes  int64      `json:"average_simpedes" db:"average_simpedes"`
	TotalDPk         float64    `json:"total_dpk" db:"total_dpk"`
	OsDPKKupedes     float64    `json:"os_dpk_kupedes" db:"os_dpk_kupedes"`
	OsDPKBriguna     int64      `json:"os_dpk_briguna" db:"os_dpk_briguna"`
	OsDPKKur         float64    `json:"os_dpk_kur" db:"os_dpk_kur"`
	TotalNpl         float64    `json:"total_npl" db:"total_npl"`
	OsNPLKupedes     float64    `json:"os_npl_kupedes" db:"os_npl_kupedes"`
	OsNPLBriguna     int64      `json:"os_npl_briguna" db:"os_npl_briguna"`
	OsNPLKur         float64    `json:"os_npl_kur" db:"os_npl_kur"`
	RecoveryDH       float64    `json:"recovery_dh" db:"recovery_dh"`
	RealisasiPH      int64      `json:"realisasi_ph" db:"realisasi_ph"`
	TotalDebitur     int64      `json:"total_debitur" db:"total_debitur"`
	DebiturKupedes   int64      `json:"debitur_kupedes" db:"debitur_kupedes"`
	DebiturBriguna   int64      `json:"debitur_briguna" db:"debitur_briguna"`
	DebiturKur       int64      `json:"debitur_kur" db:"debitur_kur"`
	RekeningSimpedes int64      `json:"rekening_simpedes" db:"rekening_simpedes"`
	RealisasiKur     float64    `json:"realisasi_kur" db:"realisasi_kur"`
	TahunDasarOS     float64    `json:"tahun_dasar_os" db:"tahun_dasar_os"`
}
