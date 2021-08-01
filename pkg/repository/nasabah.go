package repository

import (
	"database/sql"
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

func (repo *repository) BUlkUpsertNasabah(tx *sql.Tx) error {
	querySelect := `SELECT * FROM nasabah n WHERE id=(SELECT max(id) FROM nasabah m2);`
	queryInsert := `INSERT INTO nasabah (nama, no_ktp, nomor_rekening, cif_no, saldo, no_telepon, nama_ibu_kandung, kode_pos_identitas, alamat_identitas, kode_pos_kantor, alamat_kantor, description, created_at, updated_at)
						SELECT u.` + "`Nama Debitur`" + `,  u.Lancar , u.` + "`Nomor rekening`" + `, u.CIFNO, 0, u.Lancar, u.Lancar, u.` + "`Kode Pos Identitas`" + `, u.` + "`Alamat Identitas`" + `, u.` + "`Kode Pos Kantor`" + `, u.` + "`Alamat Kantor`" + `, u.Lancar, NOW(), NOW() FROM upload u 
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
