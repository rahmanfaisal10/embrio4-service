package response

import "time"

type MiniDashboardsResponse struct {
	Periode                  time.Time `json:"periode" db:"periode"`
	PosisiDPK                float64   `json:"posisi_dpk" db:"posisi_dpk"`
	CountPosisiDPK           int64     `json:"count_posisi_dpk" db:"count_posisi_dpk"`
	PercenPosisiDPK          float64   `json:"percen_posisi_dpk" db:"percen_posisi_dpk"`
	DpkJanjiStor             float64   `json:"dpk_janji_stor" db:"dpk_janji_stor"`
	CountDpkJanjiStor        int64     `json:"count_dpk_janji_stor" db:"count_dpk_janji_stor"`
	PercenDpkJanjiStor       float64   `json:"percen_dpk_janji_stor" db:"percen_dpk_janji_stor"`
	DpkTidakBayar            float64   `json:"dpk_tidak_bayar" db:"dpk_tidak_bayar"`
	CountDpkTidakBayar       int64     `json:"count_dpk_tidak_bayar" db:"count_dpk_tidak_bayar"`
	PercenDpkTidakBayar      float64   `json:"percen_dpk_tidak_bayar" db:"percen_dpk_tidak_bayar"`
	DpkObAgf                 float64   `json:"dpk_ob_agf" db:"dpk_ob_ogf"`
	CountDpkObAgf            int64     `json:"count_dpk_ob_ogf" db:"count_dpk_ob_ogf"`
	PercenDpkObAgf           float64   `json:"percen_dpk_ob_ogf" db:"percen_dpk_ob_ogf"`
	DpkBelumJatuhTempo       float64   `json:"dpk_belum_jatuh_tempo" db:"dpk_belum_jatuh_tempo"`
	CountDpkBelumJatuhTempo  int64     `json:"count_dpk_belum_jatuh_tempo" db:"count_dpk_belum_jatuh_tempo"`
	PercenDpkBelumJatuhTempo float64   `json:"percen_dpk_belum_jatuh_tempo" db:"percen_dpk_belum_jatuh_tempo"`
	PrognosaDpk              float64   `json:"prognosa_dpk" db:"prognosa_dpk"`
	CountPrognosaDpk         int64     `json:"count_prognosa_dpk" db:"count_prognosa_dpk"`
	PercenPrognosaDpk        float64   `json:"percen_prognosa_dpk" db:"percen_prognosa_dpk"`
}

type GetAllJatuhTempoResponse struct {
	DpkTotal   float64 `json:"dpk_total" db:"dpk_total"`
	Count      int64   `json:"count" db:"count"`
	CountTotal int64   `json:"count_total" db:"count_total"`
	Status     string  `json:"status"`
}
