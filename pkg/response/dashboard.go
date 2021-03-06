package response

import "time"

type ViewDashboard struct {
	Periode                    time.Time `json:"periode" db:"periode"`
	TargetOs                   float64   `json:"target_os" db:"target_os"`
	PencapaianOs               float64   `json:"pencapaian_os" db:"pencapaian_os"`
	MinimalDeltaNilai4         float64   `json:"minimal_delta_nilai_4" db:"minimal_delta_nilai_4"`
	AngsuranPokokBelumMasuk    float64   `json:"angsuran_pokok_belum_masuk" db:"angsuran_pokok_belum_masuk"`
	PencapaianRealisasi        float64   `json:"pencapaian_realisasi" db:"pencapaian_realisasi"`
	SisaSuplesi                float64   `json:"sisa_suplesi" db:"sisa_suplesi"`
	RincianLunasHutang         float64   `json:"rincian_lunas_hutang" db:"rincian_lunas_hutang"`
	DhTotal                    float64   `json:"dh_total" db:"dh_total"`
	DpkPosisi                  float64   `json:"dpk_posisi" db:"dpk_posisi"`
	CountDpkPosisi             int64     `json:"count_dpk_posisi" db:"count_dpk_posisi"`
	DpkBaru                    float64   `json:"dpk_baru" db:"dpk_baru"`
	CountDpkBaru               int64     `json:"count_dpk_baru" db:"count_dpk_baru"`
	Dpk1                       float64   `json:"dpk_1" db:"dpk_1"`
	CountDpk1                  int64     `json:"count_dpk_1" db:"count_dpk_1"`
	Dpk2                       float64   `json:"dpk_2" db:"dpk_2"`
	CountDpk2                  int64     `json:"count_dpk_2" db:"count_dpk_2"`
	Dpk3                       float64   `json:"dpk_3" db:"dpk_3"`
	CountDpk3                  int64     `json:"count_dpk_3" db:"count_dpk_3"`
	DPKCalonNPL                float64   `json:"dpk_calon_npl" db:"dpk_calon_npl"`
	CountDpkCalonNpl           int64     `json:"count_dpk_calon_npl" db:"count_dpk_calon_npl"`
	DPKBelumRestruk            float64   `json:"dpk_blm_restruk" db:"dpk_blm_restruk"`
	CountDpkBelumRestruk       int64     `json:"count_dpk_blm_restruk" db:"count_dpk_blm_restruk"`
	LancarBelumJatuhTempo      float64   `json:"lancar_blm_jatuh_tempo" db:"lancar_blm_jatuh_tempo"`
	CountLancarBelumJatuhTempo int64     `json:"count_lancar_blm_jatuh_tempo" db:"count_lancar_blm_jatuh_tempo"`
	NplTotal                   float64   `json:"npl_total" db:"npl_total"`
	CountNplTotal              int64     `json:"count_npl_total" db:"count_npl_total"`
	NplKLBaru                  float64   `json:"npl_Kl_baru" db:"npl_Kl_baru"`
	CountNplKLBaru             int       `json:"count_npl_kl_baru" db:"count_npl_kl_baru"`
	NplKl                      float64   `json:"npl_Kl" db:"npl_Kl"`
	CountNplKL                 int64     `json:"count_npl_kl" db:"count_npl_kl"`
	NplDiragukan               float64   `json:"npl_diragukan_total" db:"npl_diragukan_total"`
	CountNplDiragukan          int       `json:"count_npl_diragukan_total" db:"count_npl_diragukan_total"`
	NplMacet                   float64   `json:"npl_macet" db:"npl_macet"`
	CountNplMacet              int64     `json:"count_npl_macet" db:"count_npl_macet"`
	NplBelumRestruk            float64   `json:"npl_belum_restruk" db:"npl_belum_restruk"`
	CountNplBelumRestruk       int64     `json:"count_npl_belum_restruk" db:"count_npl_belum_restruk"`
}
