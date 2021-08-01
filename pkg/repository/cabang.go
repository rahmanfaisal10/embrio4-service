package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func (repo *repository) BulkUpsertCabang(request []*model.Cabang, tx *sqlx.Tx) ([]*model.Cabang, error) {
	querySelect := `SELECT * FROM cabang c WHERE id=(SELECT max(id) FROM cabang c2);`
	queryInsert := `INSERT INTO cabang (kode, nama, alamat, description, created_at, updated_at) VALUES %s ON DUPLICATE KEY UPDATE updated_at = NOW();`
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
		return nil, err
	}

	for _, v := range request {
		valueStrings := "(?, ?, ?, ?, NOW(), NOW())"

		valueArgss := []interface{}{
			v.Kode,
			v.Nama,
			v.Alamat,
			v.Description,
		}

		query := fmt.Sprintf(queryInsert, valueStrings)
		_, err := tx.Exec(query, valueArgss...)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, err
		}
	}

	return request, nil
}
