package repository

import (
	"database/sql"
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

func (repo *repository) BulkUpsertUnit(tx *sql.Tx) error {
	querySelect := `SELECT * FROM unit u WHERE id=(SELECT max(id) FROM unit u2);`
	queryInsert := `INSERT INTO unit (id_cabang, kode, nama, alamat, description, created_at, updated_at)
	SELECT c.id, u.` + "`Main Branch`" + ` , u.Lancar , u.Lancar , u.Lancar , NOW(), NOW() from upload u join cabang c on u.Branch = c.kode 
	ON DUPLICATE KEY UPDATE updated_at = NOW();`
	queryReset := `ALTER TABLE unit AUTO_INCREMENT = %d`

	//select last id
	unit := new(model.Unit)
	tx.QueryRow(querySelect).Scan(&unit.ID, &unit.IDCabang, &unit.Kode, &unit.Nama, &unit.Alamat, &unit.Description, &unit.CreatedAt, &unit.UpdatedAt)

	//reset auto increment
	smtf := fmt.Sprintf(queryReset, unit.ID+1)
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
