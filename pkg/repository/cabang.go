package repository

import (
	"database/sql"
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

func (repo *repository) BulkUpsertCabang(tx *sql.Tx) error {
	querySelect := `SELECT * FROM cabang c WHERE id=(SELECT max(id) FROM cabang c2);`
	queryInsert := `INSERT INTO cabang 
					(kode, nama, alamat, description, created_at, updated_at) 
					select u.Branch, u.Lancar , u.Lancar , u.Lancar , NOW(), NOW() from upload u
					ON DUPLICATE KEY UPDATE 
					nama = value(nama),
					alamat = value(alamat),
					description = value(description),
					updated_at = NOW();`

	queryReset := `ALTER TABLE cabang AUTO_INCREMENT = %d`

	//select last id
	cabang := new(model.Cabang)
	tx.QueryRow(querySelect).Scan(&cabang.ID, &cabang.Kode, &cabang.Nama, &cabang.Alamat, &cabang.Description, &cabang.CreatedAt, &cabang.UpdatedAt)

	//reset auto increment
	smtf := fmt.Sprintf(queryReset, cabang.ID+1)
	_, err := tx.Exec(smtf)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(queryInsert)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	return nil
}
