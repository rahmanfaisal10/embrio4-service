package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

var (
	AMOUNTFIELDMANTRI = 7
)

func (repo *repository) BUlkUpsertMantri(request []*model.Mantri, tx *sqlx.Tx) ([]*model.Mantri, error) {
	querySelect := `SELECT * FROM mantri m WHERE id=(SELECT max(id) FROM mantri m2);`
	queryInsert := `INSERT INTO mantri (id_unit, kode, nama, alamat, description, created_at, updated_at) VALUES %s 
					ON DUPLICATE KEY UPDATE 
						nama = value(nama),
						alamat = value(alamat),
						description = value(description),
						updated_at = NOW();`
	queryReset := `ALTER TABLE mantri AUTO_INCREMENT = %d`
	queryUnit := `SELECT * FROM unit m WHERE m.kode = %s`

	//select last id
	mantri := new(model.Mantri)
	tx.QueryRow(querySelect).Scan(&mantri.ID, &mantri.IDUnit, &mantri.Kode, &mantri.Nama, &mantri.Alamat, &mantri.Description, &mantri.CreatedAt, &mantri.UpdatedAt)

	//reset auto increment
	smtfMantri := fmt.Sprintf(queryReset, mantri.ID+1)
	_, err := tx.Exec(smtfMantri)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return nil, err
	}

	//asynchrone proses in insert query
	valueStrings := []string{}
	valueArgs := []interface{}{}

	for _, v := range request {
		unit := new(model.Unit)
		err := tx.QueryRow(fmt.Sprintf(queryUnit, v.Description)).Scan(&unit.ID, &unit.IDCabang, &unit.Kode, &unit.Nama, &unit.Alamat, &unit.Description, &unit.CreatedAt, &unit.UpdatedAt)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		valueStrings := "(?, ?, ?, ?, ?, NOW(), NOW())"

		valurArgs := []interface{}{
			unit.ID,
			v.Kode,
			v.Nama,
			v.Alamat,
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

	query := fmt.Sprintf(queryInsert, strings.Join(valueStrings, ","))

	_, err = tx.Exec(query, valueArgs...)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return nil, err
	}

	return request, nil
}
