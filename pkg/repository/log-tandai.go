package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

func (repo *repository) InsertLogTandai(request *model.LogTandai) error {
	queryInsert := `INSERT INTO embrio4.log_tandai
	(pn_pengguna, nama_pengguna, periode, cifno, nomor_rekening_simpanan, nomor_rekening_pinjaman, status, tgl_janji_setor, created_at, updated_at)
	VALUES %s;`

	querySelect := `SELECT id FROM embrio4.log_tandai lt WHERE id=(SELECT max(id) FROM embrio4.log_tandai lt2);`
	queryReset := `ALTER TABLE embrio4.log_tandai AUTO_INCREMENT = %d`

	//select last id
	logHistory := new(model.LogTandai)
	_ = repo.db.QueryRow(querySelect).Scan(logHistory.ID)

	// reset auto increment
	smtfLog := fmt.Sprintf(queryReset, logHistory.ID+1)
	_, err := repo.db.Exec(smtfLog)
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

	query := fmt.Sprintf(queryInsert, " (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
	_, err = tx.Exec(query, request.PnPengguna, request.NamaPenguna, request.Periode, request.CifNo, request.NomorRekeningSimpanan, request.NomorRekeningPinjaman, request.Status, request.TglJanjiSetor)
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
