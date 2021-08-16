package repository

import (
	"database/sql"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/labstack/gommon/log"
)

func (repo *repository) KondisiOS(mantri string, tx sql.Tx) ([]*response.TotalOS, error) {
	total := make([]*response.TotalOS, 0)
	query := `SELECT 
				u.periode, 
				u.pn_pengelola,
				COALESCE(ts.total_os,0) as total_os,
				COALESCE(u3.pencapaian_os,0) as os,
				COALESCE(((ts.total_os * 116/100)-u3.OS)/(12-MONTH(CURDATE())+1),0) as minimal_delta_nilai_4,
				coalesce(u3.OS/(((ts.total_os - ts.tahun_dasar_os)/12 * MONTH(u.Periode))+ts.tahun_dasar_os)*100,0) as percentage_pencapaian_os_periode,
				(case when u.next_pmt_date >= u.Periode then u.Plafond / CAST(TRIM(BOTH 'M' FROM u.jangka_waktu) as int) else 0 end) as angsuran_pokok_belum_masuk,
				(case when  DATEDIFF(last_day(NOW()), u.next_pmt_date) > 271 then u3.os else 0 end) as calon_dh_bulan_ini
			FROM target_smk ts 
			join upload u 
				on u.pn_pengelola = ts.pn
			join (
				SELECT 
					u2.periode,
					u2.pn_pengelola,
					SUM(u2.kolektibilitas_lancar+u2.kolektebilitas_dpk+u2.kolektebilitas_kurang_lancar+u2.kolektebilitas_diragukan+u2.kolektebilitas_macet) as pencapaian_os
				FROM upload u2
				WHERE u2.pn_pengelola is NOT NULL 
				group by u2.pn_pengelola, u2.periode
			)u3
				on u3.pn_pengelola = u.pn_pengelola
			where u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 1 month)) FROM upload u2) and ts.pn = ?;`

	rows, err := tx.Query(query, mantri)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		totalOs := new(response.TotalOS)
		if rows.Scan(
			&totalOs.Periode,
			&totalOs.PNPengelola,
			&totalOs.TargetOs,
			&totalOs.Os,
			&totalOs.MinimalDeltaNilai4,
			&totalOs.CalonDHBulanIni,
		); err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, err
		}
		total = append(total, totalOs)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Error(err)
		return nil, err
	}

	return total, nil
}

func (repo *repository) GetAllUpload() ([]*model.Upload, error) {
	upload := make([]*model.Upload, 0)
	query := `SELECT * FROM embrio4.upload u 
			WHERE u.periode = (SELECT MAX(u2.periode) FROM upload u2)`

	err := repo.db.Select(&upload, query)
	if err != nil {
		return nil, err
	}

	return upload, nil
}

func (repo *repository) UploadBulanSeblumnya(cif string) (*float64, error) {
	sisaSuplesi := new(float64)
	query := `SELECT
				(u2.kolektibilitas_lancar + u2.kolektibilitas_dpk +u2.kolektibilitas_kurang_lancar +u2.kolektibilitas_diragukan +u2.kolektibilitas_macet) as sisa_os
			FROM upload u2 
			WHERE
				u2.Periode = last_day(DATE_SUB((SELECT max(u.periode) from upload u), interval 1 month)) and u2.cif_no = ?`

	err := repo.db.Get(sisaSuplesi, query, cif)
	if err != nil {
		if err == sql.ErrNoRows {
			return sisaSuplesi, nil
		}
		return nil, err
	}

	return sisaSuplesi, nil
}

func (repo *repository) RincianLunasHutang(cif string) (*float64, error) {
	lunasHutang := new(float64)
	query := `SELECT
					u.sisa_os as sisa_lunas_hutang
				FROM (SELECT
					u2.pn_pengelola,
					u2.tgl_realisasi,
					u2.cif_no,
					(u2.kolektibilitas_lancar + u2.kolektibilitas_dpk +u2.kolektibilitas_kurang_lancar +u2.kolektibilitas_diragukan +u2.kolektibilitas_macet) as sisa_os
				FROM upload u2 
				WHERE
					u2.periode = last_day(DATE_SUB((SELECT max(u.periode) from upload u), interval 1 month))) u
				WHERE NOT EXISTS 
				(
					SELECT u3.periode, u3.cif_no from upload u3 
					WHERE 
					u3.Periode = (SELECT max(upload.periode) from upload) AND 
					(u3.cif_no = u.cif_no) and u3.cif_no = ?
				);`

	err := repo.db.Get(lunasHutang, query, cif)
	if err != nil {
		if err == sql.ErrNoRows {
			return lunasHutang, nil
		}
		return nil, err
	}

	return lunasHutang, nil
}
