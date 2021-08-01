package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func (repo *repository) BUlkUpsertNasabah(request []*model.Nasabah, tx *sqlx.Tx) ([]*model.Nasabah, error) {
	querySelect := `SELECT * FROM nasabah n WHERE id=(SELECT max(id) FROM nasabah m2);`
	queryInsert := `INSERT INTO nasabah (nama, no_ktp, nomor_rekening, cif_no, saldo, no_telepon, nama_ibu_kandung, kode_pos_identitas, alamat_identitas, kode_pos_kantor, alamat_kantor, description, created_at, updated_at)
					VALUES %s 
					ON DUPLICATE KEY UPDATE 
						nama = value(nama),
						saldo = value(saldo),
						no_telepon = value(no_telepon),
						nama_ibu_kandung = value(nama_ibu_kandung),
						kode_pos_identitas = value(kode_pos_identitas),
						alamat_identitas = value(alamat_identitas),
						kode_pos_kantor = value(kode_pos_kantor),
						alamat_kantor = value(alamat_kantor),
						updated_at = NOW();`
	queryReset := `ALTER TABLE nasabah AUTO_INCREMENT = %d`

	//select last id
	nasabah := new(model.Nasabah)
	tx.QueryRow(querySelect).Scan(
		&nasabah.ID,
		&nasabah.Nama,
		&nasabah.NoKtp,
		&nasabah.NomorRekening,
		&nasabah.CifNo,
		&nasabah.Saldo,
		&nasabah.NoTelepon,
		&nasabah.NamaIbuKandung,
		&nasabah.KodePosIdentitas,
		&nasabah.AlamatIdentitas,
		&nasabah.KodePosKantor,
		&nasabah.AlamatKantor,
		&nasabah.Description,
		&nasabah.CreatedAt,
		&nasabah.UpdatedAt,
	)

	//reset auto increment
	smtfNasabah := fmt.Sprintf(queryReset, nasabah.ID+1)
	_, err := tx.Exec(smtfNasabah)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return nil, err
	}

	for _, w := range request {

		valueStrings := "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())"

		valueArgs := []interface{}{
			w.Nama,
			w.NoKtp,
			w.NomorRekening,
			w.CifNo,
			w.Saldo,
			w.NoTelepon,
			w.NamaIbuKandung,
			w.KodePosIdentitas,
			w.AlamatIdentitas,
			w.KodePosKantor,
			w.AlamatKantor,
			"",
		}

		query := fmt.Sprintf(queryInsert, valueStrings)

		_, err = tx.Exec(query, valueArgs...)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, err
		}
	}

	return request, nil
}
