package response

import "time"

type TotalOS struct {
	Periode                       time.Time `json:"periode" db:"periode"`
	TargetOs                      float64   `json:"target_os" db:"target_os"`
	Os                            float64   `json:"os" db:"os"`
	PercentageOSAkhirTahun        float64   `json:"percentage_os_akhir_tahun" db:"percentage_os_akhir_tahun"`
	MinimalDeltaNilai4            float64   `json:"minimal_delta_nilai_4" db:"minimal_delta_nilai_4"`
	PercentagePencapaianOSPeriode float64   `json:"percentage_pencapaian_os_periode" db:"percentage_pencapaian_os_periode"`
	AngsuranPokokBelumMasuk       float64   `json:"angsuran_pokok_belum_masuk" db:"angsuran_pokok_belum_masuk"`
	CalonDHBulanIni               float64   `json:"calon_dh_bulan_ini" db:"calon_dh_bulan_ini"`
}
