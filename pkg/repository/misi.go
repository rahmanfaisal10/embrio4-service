package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func (repo *repository) BUlkUpsertMisi(request []*model.Misi, tx *sqlx.Tx) error {
	querySelect := `SELECT * FROM misi m WHERE id=(SELECT max(id) FROM misi m2);`
	queryInsert := `INSERT INTO misi (id_mantri, kode, nama, description, created_at, updated_at) VALUES %s 
					ON DUPLICATE KEY UPDATE 
						nama = value(nama),
						description = value(description),
						updated_at = NOW();`
	queryReset := `ALTER TABLE misi AUTO_INCREMENT = %d`
	queryMantri := `SELECT * FROM mantri m WHERE m.kode = %s`

	//select last id
	misi := new(model.Misi)
	tx.QueryRow(querySelect).Scan(&misi.ID, &misi.IDMantri, &misi.Kode, &misi.Nama, &misi.Description, &misi.CreatedAt, &misi.UpdatedAt)

	//reset auto increment
	smtfMisi := fmt.Sprintf(queryReset, misi.ID+1)
	_, err := tx.Exec(smtfMisi)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	valueStrings := []string{}
	valueArgs := []interface{}{}

	for _, v := range request {

		mantri := new(model.Mantri)
		err := tx.QueryRow(fmt.Sprintf(queryMantri, v.Description)).Scan(&mantri.ID, &mantri.IDUnit, &mantri.Kode, &mantri.Nama, &mantri.Alamat, &mantri.Description, &mantri.CreatedAt, &mantri.UpdatedAt)
		if err != nil {
			log.Error(err)
			return err
		}
		valueStrings = append(valueStrings, "(?, ?, ?, ?, NOW(), NOW())")

		valueArgs = append(valueArgs, mantri.ID)
		valueArgs = append(valueArgs, v.Kode)
		valueArgs = append(valueArgs, v.Nama)
		valueArgs = append(valueArgs, "")
	}

	queryInsert = fmt.Sprintf(queryInsert, strings.Join(valueStrings, ","))
	_, err = tx.Exec(queryInsert, valueArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
