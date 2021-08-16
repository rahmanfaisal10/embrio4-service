package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

func (repo *repository) UploadRepository(request []*model.Upload) error {
	queryInsert := `INSERT INTO embrio4.upload
	(periode, branch, currency, nama_ao, ln_type, nomor_rekening, nama_debitur, plafond, next_pmt_date, next_int_pmt_date, rate, tgl_menunggak, tgl_realisasi, tgl_jatuh_tempo, jangka_waktu, flag_restruk, cif_no, kolektibilitas_lancar, kolektibilitas_dpk, kolektibilitas_kurang_lancar, kolektibilitas_diragukan, kolektibilitas_macet, tunggakan_pokok, tunggakan_bunga, tunggakan_pinalty, pn_pengelola, nama_pengelola, code, description, kol_adk, avg_os_harian, kecamatan_tempat_tinggal, kelurahan_tempat_tinggal, kode_pos_tempat_tinggal, kecamatan_tempat_usaha, kelurahan_tempat_usaha, kode_pos_tempat_usaha)
	VALUES %s`

	querySelect := `SELECT id FROM upload m WHERE id=(SELECT max(id) FROM upload m2);`
	queryReset := `ALTER TABLE upload AUTO_INCREMENT = %d`

	//select last id
	upload := new(model.Upload)
	_ = repo.db.QueryRow(querySelect).Scan(&upload.ID)

	//reset auto increment
	smtfUpload := fmt.Sprintf(queryReset, upload.ID+1)
	_, err := repo.db.Exec(smtfUpload)
	if err != nil {
		return err
	}

	//config transaction mode
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	//bulk insert into table upload

	query := fmt.Sprintf(queryInsert, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	for _, row := range request {
		vals := []interface{}{row.Periode, row.Branch, row.Currency, row.NamaAO, row.LNType, row.NomorRekening, row.NamaDebitur, row.Plafond, row.NextPmtDate, row.NextIntPmtDate, row.Rate, row.TglMenunggak, row.TglRealisasi, row.TglJatuhTempo, row.JangkaWaktu, row.FlagRestruk, row.CIFNO, row.KolektibilitasLancar, row.KolektibilitasDPK, row.KolektibilitasKurangLancar, row.KolektibilitasDiragukan, row.KolektibilitasMacet, row.TunggakanPokok, row.TunggakanBunga, row.TunggakanPinalty, row.PNPengelola, row.NamaPengelola, row.Code, row.Description, row.KolADK, row.AvgOsHarian, row.KecamatanTempatTinggal, row.KelurahanTempatTinggal, row.KodePosTempatTinggal, row.KecamatanTempatUsaha, row.KelurahanTempatUsaha, row.KodePosTempatUsaha}
		_, err = tx.Exec(query, vals...)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			return err
		}
	}

	err = repo.BUlkUpsertMantri(tx)
	if err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}

	return nil
}
