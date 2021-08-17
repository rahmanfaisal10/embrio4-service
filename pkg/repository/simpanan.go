package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

func (repo *repository) InsertSimpanan(request []*model.Simpanan) error {
	queryInsert := `INSERT INTO embrio4.simpanan
	(periode, uker_code, curr_code, curr_desc, account_number, ciff_no, short_name, officer_name, dlt, Open_dt, balance, available_balance, int_credit, accrued_int, average_balance, prod_code, pn_pengelola, nama_pengelola, kecamatan_tempat_tinggal, kelurahan_tempat_tinggal, kode_pos_tempat_tinggal, kecamatan_tempat_usaha, kelurahan_tempat_usaha, kode_pos_tempat_usaha)
	VALUES %s;`

	querySelect := `SELECT id FROM simpanan m WHERE id=(SELECT max(id) FROM simpanan m2);`
	queryReset := `ALTER TABLE simpanan AUTO_INCREMENT = %d`

	//select last id
	simpanan := new(model.Simpanan)
	_ = repo.db.QueryRow(querySelect).Scan(&simpanan.ID)

	//reset auto increment
	smtfSimpanan := fmt.Sprintf(queryReset, simpanan.ID+1)
	_, err := repo.db.Exec(smtfSimpanan)
	if err != nil {
		log.Error(err)
		return err
	}

	//config transaction mode
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}

	query := fmt.Sprintf(queryInsert, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	for _, v := range request {
		vals := []interface{}{v.Periode, v.UkerCode, v.CurrCode, v.CUrrDesc, v.AccountNumber, v.CifNo, v.ShortName, v.OfficerName, v.Dlt, v.OpenDT, v.Balance, v.AvailableBalance, v.IntCredit, v.AccruedInt, v.AverageBalance, v.ProdCode, v.PnPengelola, v.NamaPengelola, v.KecamatanTempatTinggal, v.KelurahanTempatTinggal, v.KodePosTempatTinggal, v.KecamatanTempatUsaha, v.KelurahanTempatUsaha, v.KodePosTempatUsaha}
		_, err = tx.Exec(query, vals...)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}

	return nil

}
