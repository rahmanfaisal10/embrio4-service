package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func (repo *repository) BulkUpsertUnit(request []*model.Unit, tx *sqlx.Tx) ([]*model.Unit, error) {
	querySelect := `SELECT * FROM unit u WHERE id=(SELECT max(id) FROM unit u2);`
	queryInsert := `INSERT INTO unit (id_cabang, kode, nama, alamat, description, created_at, updated_at) VALUES %s ON DUPLICATE KEY UPDATE updated_at = NOW();`
	queryReset := `ALTER TABLE unit AUTO_INCREMENT = %d`
	queryCabang := `SELECT * FROM cabang c WHERE c.kode = %s`

	//select last id
	unit := new(model.Unit)
	tx.QueryRow(querySelect).Scan(&unit.ID, &unit.IDCabang, &unit.Kode, &unit.Nama, &unit.Alamat, &unit.Description, &unit.CreatedAt, &unit.UpdatedAt)

	//reset auto increment
	smtf := fmt.Sprintf(queryReset, unit.ID+1)
	_, err := tx.Exec(smtf)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return nil, err
	}

	//insert cabang
	for _, w := range request {
		cabang := new(model.Cabang)
		err := tx.QueryRow(fmt.Sprintf(queryCabang, w.Description)).Scan(&cabang.ID, &cabang.Kode, &cabang.Nama, &cabang.Alamat, &cabang.Description, &cabang.CreatedAt, &cabang.UpdatedAt)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, err
		}
		valueStrings := "(?, ?, ?, ?, ?, NOW(), NOW())"

		valurArgs := []interface{}{
			cabang.ID,
			w.Kode,
			w.Nama,
			w.Alamat,
			"",
		}

		query := fmt.Sprintf(queryInsert, valueStrings)
		_, err = tx.Exec(query, valurArgs...)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, err
		}
	}

	return request, nil
}
