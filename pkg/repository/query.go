package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/response"
)

func (repo *repository) QueryTotalOS(mantri string) (totalOS response.TotalOS, err error) {
	query := `select DISTINCT
				u.Periode as periode,
				coalesce(ts.total_os, 0) as "target_os",
				coalesce(u3.os,0) as os,
				coalesce(u3.OS/total_os*100,0) as "percentage_os_akhir_tahun",
				coalesce(((ts.total_os * 116/100)-u3.OS)/(12-MONTH(CURDATE())+1),0) as minimal_delta_nilai_4,
				coalesce(u3.OS/(((ts.total_os - ts.tahun_dasar_os)/12 * MONTH(u.Periode))+ts.tahun_dasar_os)*100,0) as percentage_pencapaian_os_periode,
				(case when u.` + "`Next Pmt Date`" + ` >= u.Periode then u.Plafond / CAST(TRIM(BOTH 'M' FROM u.` + "`Jangka Waktu`" + `) as int) else 0 end) as angsuran_pokok_belum_masuk,
				(case when  DATEDIFF(last_day(NOW()), u.` + "`Next Pmt Date`" + `) > 271 then u3.os else 0 end) as calon_dh_bulan_ini
				
			from target_smk ts
			join upload u 
				on u.` + "`PN   PENGELOLA`" + ` = ts.pn
			join
				(
					select 
						u2.` + "`PN   PENGELOLA`" + ` as pn,
						SUM( u2.` + "`Kolektibilitas Lancar`" + ` +	u2.` + "`Kolektibilitas DPK`" + ` + u2.` + "`Kolektibilitas Kurang Lancar`" + ` +	u2.` + "`Kolektibilitas Diragukan`" + ` + 	u2.` + "`Kolektibilitas Macet`" + ` ) as "os"
					FROM upload u2
					group by u2 .` + "`PN   PENGELOLA`" + `
				) u3
				on u3.pn = u.` + "`PN   PENGELOLA`" + `
			where u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 1 month)) FROM upload u2) and ts.pn = %s;`

	err = repo.db.Get(&totalOS, fmt.Sprintf(query, mantri))
	if err != nil {
		return response.TotalOS{}, err
	}

	return

}

func (repo *repository) QuerySisaSuplesi(mantri string) (sisaSuplesi response.SisaSuplesiResponse, err error) {
	query := `
			SELECT 
				COUNT(*) as count , 
				u.Periode as periode, 
				bulan_sebelumnya.Periode as periode_bulan_sebelumnya, 
				u.CIFNO , 
				u.` + "`Nomor rekening`" + ` as nomor_rekening , 
				bulan_sebelumnya.` + "`Nomor rekening`" + ` as nomor_rekening_sebelumnya,  
				SUM(bulan_sebelumnya.sisa_os) as sisa_suplesi
			FROM upload u 
			JOIN
				(SELECT
					u2.` + "`PN   PENGELOLA`" + `,
					u2.` + "`Tgl Realisasi`" + `,
					u2.CIFNO,
					u2.Periode ,
					u2.` + "`Nomor rekening`" + `,
					(u2.` + "`Kolektibilitas Lancar`" + ` + u2.` + "`Kolektibilitas DPK`" + ` +u2.` + "`Kolektibilitas Kurang Lancar`" + ` +u2.` + "`Kolektibilitas Diragukan`" + ` +u2.` + "`Kolektibilitas Macet`" + `) as sisa_os
				FROM upload u2 
				WHERE
					u2.Periode = last_day(DATE_SUB((SELECT max(u.periode) from upload u), interval 1 month)) ) bulan_sebelumnya
				ON
					u.CIFNO = bulan_sebelumnya.CIFNO
			WHERE 
				u.Periode = (SELECT max(m.periode) from upload m) and u.` + "`PN   PENGELOLA`" + ` = %s
			GROUP by u.` + "`PN   PENGELOLA`" + `;`

	err = repo.db.Get(&sisaSuplesi, fmt.Sprintf(query, mantri))
	if err != nil {
		return response.SisaSuplesiResponse{}, err
	}
	return
}

func (repo *repository) QueryRincianLunasHutang(mantri string) (sisaLunasHutang response.SisaLunasHutangResponse, err error) {
	query := `
	SELECT 	COUNT(*) as count, 
		coalesce(SUM(u.sisa_os),0) as sisa_lunas_hutang, 
		coalesce(u.` + "`PN   PENGELOLA`" + `,'') as pn_pengelola,
		u.` + "`Tgl Realisasi`" + ` as tgl_realisasi,
		coalesce(u.CIFNO,'') as cif_no,
		u.Periode as periode,
		coalesce(u.Branch,'') as branch,
		coalesce(u.` + "`Nomor rekening`" + `,'') as nomor_rekening
	FROM (SELECT
		u2.` + "`PN   PENGELOLA`" + `,
		u2.` + "`Tgl Realisasi`" + `,
		u2.CIFNO,
		u2.Branch,
		u2.Periode,
		u2.` + "`Nomor rekening`" + `,
		u2.` + "`Nama Debitur`" + `,
		(u2.` + "`Kolektibilitas Lancar`" + ` + u2.` + "`Kolektibilitas DPK`" + ` +u2.` + "`Kolektibilitas Kurang Lancar`" + ` +u2.` + "`Kolektibilitas Diragukan`" + ` +u2.` + "`Kolektibilitas Macet`" + `) as sisa_os
	FROM upload u2 
	WHERE
		u2.Periode = last_day(DATE_SUB((SELECT max(u.periode) from upload u), interval 1 month))) u
	WHERE NOT EXISTS 
	(
	SELECT u3.Periode, u3.CIFNO from upload u3 
	WHERE 
	u3.Periode = (SELECT max(upload.periode) from upload) AND 
	(u3.CIFNO = u.CIFNO)
	)
	and  u.` + "`PN   PENGELOLA`" + ` = %s;`

	err = repo.db.Get(&sisaLunasHutang, fmt.Sprintf(query, mantri))
	if err != nil {
		return response.SisaLunasHutangResponse{}, err
	}
	return
}

func (repo *repository) PencapaianRealisasi(mantri string) (pencapaianRealisasi float64, err error) {
	query := `SELECT SUM(u.Plafond) as pencapaian_realisasi FROM upload u 
				where MID(u.` + "`Tgl Realisasi`" + `, 1, 7) = MID((SELECT MAX(u2.Periode) from upload u2), 1, 7) 
				and u.` + "`PN   PENGELOLA`" + ` = ?`

	err = repo.db.Get(&pencapaianRealisasi, query, mantri)
	if err != nil {
		return 0, err
	}
	return
}

// func (repo *repository) DPKBersaldoSimpanan(mantri string) {
// 	query := `SELECT COUNT(*)as count, SUM(s.`available balance`) as nomor_rekening_simpanan,SUM(u.`Kolektibilitas DPK`) as os_DPK_saldo_simpanan FROM upload u
// 	JOIN simpanan s ON u.CIFNO = s.`ciff no`
// 	WHERE u.`PN   PENGELOLA` = "00213841" and u.`Kolektibilitas DPK` != 0 and u.`Tunggakan Pokok` +u.`Tunggakan Bunga` < s.`available balance`
// 	group BY u.`Nomor rekening``
// 	return
// }

// func (repo *repository) PosisiDPK(mantri string) {
// 	query := `SELECT COUNT(*), SUM(up.`Kolektibilitas DPK`) from upload up
// 	where up.`PN   PENGELOLA` = "00212710" and up.Periode = (SELECT MAX(az.Periode) from upload az) and up.`Kolektibilitas DPK` != 0`
// }

// func (repo *repository) DPKBaru(mantri string) {
// 	query := `SELECT COUNT(*) as count, SUM(up.`Kolektibilitas DPK`) as dpk_baru from upload up
// 	join
// 	 (
// 		 SELECT u.`Nomor rekening`, u.Periode, u.`Kolektibilitas DPK`,u.`Kolektibilitas Lancar`, u.CIFNO FROM upload u
// 		WHERE u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 1 month)) FROM upload u2) and u.`Kolektibilitas Lancar` != 0
// 	 ) bu
// 	 on up.`Nomor rekening` = bu.`Nomor rekening`
// 	where up.`Kolektibilitas DPK` != 0 and bu.`Kolektibilitas Lancar` != 0 and up.`PN   PENGELOLA` = "00212710" and up.Periode = (SELECT MAX(az.Periode) from upload az)`
// }

// func (repo *repository) posisiDPK2(mantri string) {
// 	query := `SELECT COUNT(*) ,SUM(up.`Kolektibilitas DPK`) from upload up
// 	join
// 	 (
// 		 SELECT u.`Nomor rekening`, u.Periode, u.`Kolektibilitas DPK`,u.`Kolektibilitas Lancar`, u.CIFNO FROM upload u
// 		WHERE u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 1 month)) FROM upload u2) and u.`Kolektibilitas DPK` != 0
// 	 ) bu
// 	 on up.`Nomor rekening` = bu.`Nomor rekening`
// 	join
// 	 (
// 		 SELECT u.`Nomor rekening`, u.Periode, u.`Kolektibilitas DPK`,u.`Kolektibilitas Lancar`, u.CIFNO FROM upload u
// 		WHERE u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 2 month)) FROM upload u2) and u.`Kolektibilitas Lancar` != 0
// 	 ) tu
// 	 on up.`Nomor rekening` = tu.`Nomor rekening` and bu.`Nomor rekening` = tu.`Nomor rekening`
// 	where up.`Kolektibilitas DPK` != 0 and bu.`Kolektibilitas DPK` != 0 and up.`PN   PENGELOLA` = "00212710" and up.Periode = (SELECT MAX(az.Periode) from upload az)`
// }

// func (repo *reporepository) DPK3Bulan(mantri string)  {
// 	query := `SELECT COUNT(*), SUM(up.`Kolektibilitas DPK`) from upload up
// 	join
// 	 (
// 		 SELECT u.`Nomor rekening`, u.Periode, u.`Kolektibilitas DPK`,u.`Kolektibilitas Lancar`, u.CIFNO FROM upload u
// 		WHERE u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 1 month)) FROM upload u2) and u.`Kolektibilitas DPK` != 0
// 	 ) bu
// 	 on up.`Nomor rekening` = bu.`Nomor rekening`
// 	join
// 	 (
// 		 SELECT u.`Nomor rekening`, u.Periode, u.`Kolektibilitas DPK`,u.`Kolektibilitas Lancar`, u.CIFNO FROM upload u
// 		WHERE u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 2 month)) FROM upload u2) and u.`Kolektibilitas DPK` != 0
// 	 ) tu
// 	 on up.`Nomor rekening` = tu.`Nomor rekening` and bu.`Nomor rekening` = tu.`Nomor rekening`
// 	where up.`Kolektibilitas DPK` != 0 and up.`PN   PENGELOLA` = "00212710" and up.`Next Pmt Date` < (SELECT MAX(az.Periode) from upload az)  and up.Periode = (SELECT MAX(az.Periode) from upload az)
// 	`
// }

// func (repo *reporepository) DPKCalonNPL(mantri string)  {
// 	query := `SELECT COUNT(*), SUM(up.`Kolektibilitas DPK`) from upload up x
// 	join
// 	 (
// 		 SELECT u.`Nomor rekening`, u.Periode, u.`Kolektibilitas DPK`,u.`Kolektibilitas Lancar`, u.CIFNO FROM upload u
// 		WHERE u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 1 month)) FROM upload u2) and u.`Kolektibilitas DPK` != 0
// 	 ) bu
// 	 on up.`Nomor rekening` = bu.`Nomor rekening`
// 	join
// 	 (
// 		 SELECT u.`Nomor rekening`, u.Periode, u.`Kolektibilitas DPK`,u.`Kolektibilitas Lancar`, u.CIFNO FROM upload u
// 		WHERE u.Periode = (SELECT last_day(DATE_SUB(MAX(u2.Periode), INTERVAL 2 month)) FROM upload u2) and u.`Kolektibilitas DPK` != 0
// 	 ) tu
// 	 on up.`Nomor rekening` = tu.`Nomor rekening` and bu.`Nomor rekening` = tu.`Nomor rekening`
// 	where up.`Kolektibilitas DPK` != 0 and up.`PN   PENGELOLA` = "00212710" and up.`Next Pmt Date` > (SELECT MAX(az.Periode) from upload az)
// 	`
// }

// func (repo *repository) DPKBelumRestruk(mantri string)  {
// 	query:= `SELECT COUNT(*), SUM(up.`Kolektibilitas DPK`) from upload up
// 	where up.`PN   PENGELOLA` = "00212710" and up.Periode = (SELECT MAX(az.Periode) from upload az) and up.`Kolektibilitas DPK` != 0 and up.`Flag Restruk` = "N"`
// }

// func (repo *repository) LancarBelumJatuhTempo(mantri string) {
// 	query := `SELECT up.`Kolektibilitas Lancar`, up.`Next Pmt Date` , up.Periode  from upload up
// 	where up.`Next Pmt Date` < (SELECT MAX(az.Periode) from upload az) and up.Periode = (SELECT MAX(az.Periode) from upload az) and up.`Kolektibilitas Lancar` != 0
// 	and up.`PN   PENGELOLA` = "00212710"

// `
// }
