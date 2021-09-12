package repository

import (
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/labstack/gommon/log"
)

func (repo *repository) MiniDashboardRepository(mantri string) ([]*response.MiniDashboardResponse, error) {
	response := make([]*response.MiniDashboardResponse, 0)
	query := `SELECT
				d.periode, 
				SUM(d.dpk_total) as dpk_total,
				COALESCE(lt.status, '') as status,
				COUNT(*) as count
			FROM dashboard d
			LEFT JOIN log_tandai lt 
				on lt.pn_pengguna = d.Id_mantri AND 
				lt.nomor_rekening_pinjaman = d.nomor_rekening AND 
				EXTRACT(YEAR_MONTH FROM d.periode) = EXTRACT(YEAR_MONTH FROM lt.tgl_janji_setor)
			WHERE 
				d.periode = (SELECT MAX(d2.periode) FROM dashboard d2) AND
				d.dpk_total != 0 AND
				d.Id_mantri = ?
			GROUP BY lt.status`

	err := repo.db.Select(&response, query, mantri)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
}
