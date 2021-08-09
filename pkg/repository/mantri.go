package repository

import (
	"database/sql"
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

var (
	AMOUNTFIELDMANTRI = 7
)

func (repo *repository) BUlkUpsertMantri(tx *sql.Tx) error {
	querySelect := `SELECT * FROM mantri m WHERE id=(SELECT max(id) FROM mantri m2);`
	queryInsert := `INSERT INTO mantri (id_unit, kode, nama, alamat, description, created_at, updated_at)
					SELECT c.id, u.` + "`PN   PENGELOLA`" + ` , u.` + "`NAMA  PENGELOLA`" + ` , u.Lancar , u.Lancar , NOW(), NOW() from upload u join unit c on u.` + "`Main Branch`" + ` = c.kode
					WHERE u.` + "`PN   PENGELOLA`" + `IS NOT NULL and u.` + "`PN   PENGELOLA`" + ` != ''
					ON DUPLICATE KEY UPDATE 
						nama = value(nama),
						alamat = value(alamat),
						description = value(description),
						updated_at = NOW();`
	queryReset := `ALTER TABLE mantri AUTO_INCREMENT = %d`

	//select last id
	mantri := new(model.Mantri)
	tx.QueryRow(querySelect).Scan(&mantri.ID, &mantri.IDUnit, &mantri.Kode, &mantri.Nama, &mantri.Alamat, &mantri.Description, &mantri.CreatedAt, &mantri.UpdatedAt)

	//reset auto increment
	smtfMantri := fmt.Sprintf(queryReset, mantri.ID+1)
	_, err := tx.Exec(smtfMantri)
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

func (repo *repository) GetAllMantri() (mantri []*model.Mantri, err error) {
	query := `SELECT
		id, 
		IFNULL(id_unit, '') id_unit, 
		IFNULL(kode, '') kode, 
		IFNULL(nama, '') nama, 
		IFNULL(alamat, '') alamat, 
		IFNULL(description, '') description, 
		created_at, 
		updated_at
	FROM embrio4.mantri;`

	err = repo.db.Select(&mantri, query)
	if err != nil {
		return nil, err
	}
	return
}
