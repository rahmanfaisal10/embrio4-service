package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/labstack/gommon/log"
)

func (repo *repository) UploadRepository(request []*model.Upload) error {
	queryInsert := `INSERT INTO embrio4.upload
	(
		Periode,
		Region,
		` + "`Main Branch`" + `, 
		Branch,
		Currency, 
		` + "`Nama AO`" + `, 
		` + "`LN Type`" + `, 
		` + "`Nomor rekening`" + `, 
		` + "`Nama Debitur`" + `, 
		` + "`Alamat Identitas`" + `, 
		` + "`Kode Pos Identitas`" + `, 
		` + "`Alamat Kantor`" + `, 
		` + "`Kode Pos Kantor`" + `, 
		Plafond, 
		` + "`Next Pmt Date`" + `, 
		` + "`Next Int Pmt Date`" + `, 
		Rate, 
		` + "`Tgl Menunggak`" + `,
		` + "`Tgl Realisasi`" + `, 
		` + "`Tgl Jatuh tempo`" + `, 
		` + "`Jangka Waktu`" + `, 
		` + "`Flag Restruk`" + `, 
		CIFNO, 
		` + "`Kolektibilitas Lancar`" + `, 
		` + "`Kolektibilitas DPK`" + `, 
		` + "`Kolektibilitas Kurang Lancar`" + `, 
		` + "`Kolektibilitas Diragukan`" + `, 
		` + "`Kolektibilitas Macet`" + `, 
		` + "`Tunggakan Pokok`" + `, 
		` + "`Tunggakan Bunga`" + `, 
		` + "`Tunggakan Pinalty`" + `, 
		` + "`PN   PENGELOLA`" + `, 
		` + "`NAMA  PENGELOLA`" + `
	)
	VALUES %s
	ON DUPLICATE KEY UPDATE 
		Periode	= value(Periode),
		Region = value(Region),
		` + "`Main Branch`" + ` = value(` + "`Main Branch`" + `),
		Branch= value(Branch),
		Currency = value(Currency),
		` + "`Nama AO`" + ` = value(` + "`Nama AO`" + `),
		` + "`LN Type`" + ` = value(` + "`LN Type`" + `),
		` + "`Nomor rekening`" + ` = value(` + "`Nomor rekening`" + `),
		` + "`Nama Debitur`" + ` = value(` + "`Nama Debitur`" + `),
		` + "`Alamat Identitas`" + ` = value(` + "`Alamat Identitas`" + `),
		` + "`Kode Pos Identitas`" + ` = value(` + "`Kode Pos Identitas`" + `),
		` + "`Alamat Kantor`" + ` = value(` + "`Alamat Kantor`" + `),
		` + "`Kode Pos Kantor`" + ` = value(` + "`Kode Pos Kantor`" + `),
		Plafond = value(Plafond),
		` + "`Next Pmt Date`" + ` = value(` + "`Next Pmt Date`" + `),
		` + "`Next Int Pmt Date`" + ` = value(` + "`Next Int Pmt Date`" + `),
		Rate = value(Rate),
		` + "`Tgl Menunggak`" + ` = value(` + "`Tgl Menunggak`" + `),
		` + "`Tgl Realisasi`" + ` = value(` + "`Tgl Realisasi`" + `),
		` + "`Tgl Jatuh tempo`" + ` = value(` + "`Tgl Jatuh tempo`" + `),
		` + "`Jangka Waktu`" + ` = value(` + "`Jangka Waktu`" + `),
		` + "`Flag Restruk`" + ` = value(` + "`Flag Restruk`" + `),
		CIFNO = value(CIFNO),
		` + "`Kolektibilitas Lancar`" + ` = value(` + "`Kolektibilitas Lancar`" + `),
		` + "`Kolektibilitas DPK`" + ` = value(` + "`Kolektibilitas DPK`" + `),
		` + "`Kolektibilitas Kurang Lancar`" + ` = value(` + "`Kolektibilitas Kurang Lancar`" + `),
		` + "`Kolektibilitas Diragukan`" + ` = value(` + "`Kolektibilitas Diragukan`" + `),
		` + "`Kolektibilitas Macet`" + ` = value(` + "`Kolektibilitas Macet`" + `),
		` + "`Tunggakan Pokok`" + ` = value(` + "`Tunggakan Pokok`" + `),
		` + "`Tunggakan Bunga`" + ` = value(` + "`Tunggakan Bunga`" + `),
		` + "`Tunggakan Pinalty`" + ` = value(` + "`Tunggakan Pinalty`" + `),
		` + "`PN   PENGELOLA`" + ` = value(` + "`PN   PENGELOLA`" + `),
		` + "`NAMA  PENGELOLA`" + ` = value(` + "`NAMA  PENGELOLA`" + `);`

	querySelect := `SELECT id FROM upload m WHERE id=(SELECT max(id) FROM upload m2);`
	queryReset := `ALTER TABLE upload AUTO_INCREMENT = %d`

	//select last id
	upload := new(model.Upload)
	_ = repo.db.QueryRow(querySelect).Scan(&upload.ID)

	//reset auto increment
	smtfUpload := fmt.Sprintf(queryReset, upload.ID+1)
	_, err := repo.db.Exec(smtfUpload)
	if err != nil {
		return err
	}

	//config transaction mode
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	valueStrings := "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	query := fmt.Sprintf(queryInsert, valueStrings)

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	for _, v := range request {
		valueArgs := []interface{}{
			v.Periode,
			v.Region,
			v.Mainbranch,
			v.Branch,
			v.Currency,
			v.NamaAO,
			v.LNType,
			v.NomorRekening,
			v.NamaDebitur,
			v.AlamatIdentitas,
			v.KodePosIdentitas,
			v.AlamatKantor,
			v.KodePosKantor,
			v.Plafond,
			v.NextPmtDate,
			v.NextIntPmtDate,
			v.Rate,
			v.TglMenunggak,
			v.TglRealisasi,
			v.TglJatuhTempo,
			v.JangkaWaktu,
			v.FlagRestruk,
			v.CIFNO,
			v.KolektibilitasLancar,
			v.KolektibilitasDPK,
			v.KolektibilitasKurangLancar,
			v.KolektibilitasDiragukan,
			v.KolektibilitasMacet,
			v.TunggakanPokok,
			v.TunggakanBunga,
			v.TunggakanPinalty,
			v.PNPengelola,
			v.NamaPengelola,
		}

		_, err = stmt.Exec(valueArgs...)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		stmt.Close()
		log.Error(err)
		return err
	}
	stmt.Close()

	return nil
}
