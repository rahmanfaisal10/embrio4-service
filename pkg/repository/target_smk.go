package repository

import (
	"github.com/labstack/gommon/log"
)

func (repo *repository) InsertTarget() error {
	queryInsert := `INSERT INTO embrio4.target_smk
					(nama_pn, pn, total_os, tahun_dasar_os)
					SELECT DISTINCT nama_pengelola, pn_pengelola, 10000000000, 15000000000 FROM upload WHERE pn_pengelola != ''
					ON DUPLICATE KEY UPDATE tanggal = null;`

	_, err := repo.db.Exec(queryInsert)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
