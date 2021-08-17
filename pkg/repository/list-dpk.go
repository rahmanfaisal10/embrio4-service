package repository

import (
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/labstack/gommon/log"
)

func (repo *repository) ListDpkRepository(mantri string) ([]*response.ListDpkResponse, error) {
	response := make([]*response.ListDpkResponse, 0)
	query := `SELECT
				u.nama_debitur,
				u.kelurahan_tempat_tinggal as alamat,
				GROUP_CONCAT(s.available_balance) as available_balance,
				u.next_pmt_date,
				u.flag_restruk,
				u.ln_type,
				u.nomor_rekening,
				GROUP_CONCAT(s.account_number) as account_number,
				COALESCE(lt.status,'') as status 
			FROM upload u , simpanan s, dashboard d
			LEFT JOIN log_tandai lt 
				ON lt.pn_pengguna = d.Id_mantri AND lt.nomor_rekening_pinjaman = d.nomor_rekening 
			WHERE
				u.periode = (SELECT MAX(u2.periode) FROM embrio4.upload u2) AND
				d.dpk_total IS NOT NULL AND
				d.dpk_total != 0 AND 
				d.periode = (SELECT MAX(u2.periode) FROM embrio4.upload u2) AND
				d.periode = s.periode AND
				d.periode = u.periode AND
				d.nomor_rekening = u.nomor_rekening AND
				u.cif_no = s.ciff_no AND
				u.pn_pengelola = ?
			GROUP BY u.cif_no`

	err := repo.db.Select(&response, query, mantri)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
}
