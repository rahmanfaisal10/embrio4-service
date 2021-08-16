package repository

import (
	"rahmanfaisal10/embrio4-service/pkg/response"
	"rahmanfaisal10/embrio4-service/pkg/util"

	"github.com/labstack/gommon/log"
)

func (repo *repository) InsertDashboard() error {
	upload, err := repo.GetAllUpload()
	if err != nil {
		log.Error(err)
		return err
	}

	//config transaction mode
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	for _, v := range upload {
		totalOS := (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		var pencapaianRealisasi, osKupedes, kur, osKurSupermi, osGbt, totalDpk, dpkBaru, dpkSatu, dpkDua, dpkTiga, dpkFlagRestruk, lancar, NplPosisi, nplKLKecil, nplKL, nplDiragukan, nplMacet, nplRestruk float64
		thenMonth := v.Periode.AddDate(0, -1, 0)
		umurTunggakan := (v.Periode.Sub(*v.NextPmtDate).Hours() / 24)

		if v.Periode.Year() == v.TglRealisasi.Year() && v.Periode.Month() == v.TglRealisasi.Month() {
			pencapaianRealisasi = v.Plafond
		}

		sisaSuplesi, err := repo.UploadBulanSeblumnya(v.CIFNO)
		if err != nil {
			log.Error(err)
			return err
		}

		lunasHutang, err := repo.RincianLunasHutang(v.CIFNO)
		if err != nil {
			log.Error(err)
			return err
		}

		if v.LNType == "HA" || v.LNType == "HI" || v.LNType == "KJ" || v.LNType == "S1" || v.LNType == "T1" || v.LNType == "1K" || v.LNType == "1L" {
			osKupedes = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if v.LNType == "SH" || v.LNType == "SM" || v.LNType == "2F" || v.LNType == "2G" || v.LNType == "PX" {
			kur = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if v.LNType == "S" || v.LNType == "2F" || v.LNType == "2G" || v.LNType == "2I" || v.LNType == "2L" {
			osKurSupermi = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if v.LNType == "W2" || v.LNType == "KJ" {
			osGbt = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > 0 && umurTunggakan <= 90 {
			totalDpk = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > 0 && umurTunggakan <= float64(v.Periode.Day()) {
			dpkBaru = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > float64(v.Periode.Day()) && umurTunggakan <= (float64(v.Periode.Day())+float64(util.LastDay(v.Periode))) {
			dpkSatu = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > (float64(v.Periode.Day())+float64(util.LastDay(v.Periode))) && umurTunggakan <= (float64(v.Periode.Day())+float64(util.LastDay(v.Periode))+float64(util.LastDay(&thenMonth))) {
			dpkDua = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > (float64(v.Periode.Day())+float64(util.LastDay(v.Periode))+float64(util.LastDay(&thenMonth))) && umurTunggakan <= 90 {
			dpkTiga = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > 0 && umurTunggakan <= 90 && v.FlagRestruk == "N" {
			dpkFlagRestruk = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > -31 && umurTunggakan <= -1 {
			lancar = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > 90 && umurTunggakan <= 270 {
			NplPosisi = (v.KolektibilitasDPK + v.KolektibilitasDiragukan + v.KolektibilitasKurangLancar + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		if umurTunggakan > 90 && umurTunggakan <= (90+float64(v.Periode.Day())) {
			nplKLKecil = (v.KolektibilitasKurangLancar + v.KolektibilitasDiragukan + v.KolektibilitasMacet)
		}

		if v.KolADK == 3 {
			nplKL = v.KolektibilitasKurangLancar
		}

		if v.KolADK == 4 {
			nplDiragukan = v.KolektibilitasDiragukan
		}

		if v.KolADK == 5 {
			nplMacet = v.KolektibilitasMacet
		}

		if umurTunggakan > 90 && umurTunggakan <= 270 && v.FlagRestruk == "N" {
			nplRestruk = (v.KolektibilitasDiragukan + v.KolektibilitasLancar + v.KolektibilitasMacet)
		}

		query := `INSERT INTO embrio4.dashboard
		(periode, target_os, os_total, os_kupedes, os_kur, os_kur_supermi, os_gbt, sisa_suplesi, rincian_lunas_hutang, pencapaian_realisasi, dpk_total, dpk_kecil, dpk_baru, dpk_1, dpk_2, dpk_3, dpk_cnpl, dpk_blm_restruk, dpk_masih_GP, dpk_dalam_GP, lancar_blm_jth_tempo, npl_total, npl_kl_kecil, npl_kl_total, npl_diragukan_total, npl_macet, npl_belum_restruk, dh_total, ph_total, dh_pemasukan_tahunberjalan, dh_bersaldo_simpanan, dh_yg_mengangsur, simpanan_total, simpanan_topup_besar, simpanan_pengambilan_besar, simpanan_besar_tgl_lahir_bln_ini, restruk_os, restruk_dalam_grace_periode, Id_mantri, description, created_at, updated_at)
		VALUES(
			?, 
			NULL, 
			?, 
			?, 
			?, 
			?, 
			?, 
			?, 
			?, 
			?,
			?,
			0, 
			?, 
			?, 
			?, 
			?, 
			?, 
			?, 
			NULL, 
			NULL, 
			?, 
			?, 
			?, 
			?, 
			?, 
			?, 
			?, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			NULL, 
			?, 
			NULL, 
			NOW(), 
			NOW());`

		_, err = tx.Exec(query, v.Periode, totalOS, osKupedes, kur, osKurSupermi, osGbt, sisaSuplesi, lunasHutang, pencapaianRealisasi, totalDpk, dpkBaru, dpkSatu, dpkDua, dpkTiga, dpkTiga, dpkFlagRestruk, lancar, NplPosisi, nplKLKecil, nplKL, nplDiragukan, nplMacet, nplRestruk, v.PNPengelola)
		if err != nil {
			tx.Rollback()
			log.Error(err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Error(err)
		return err
	}

	return nil
}

func (repo *repository) ViewDashboard(mantri string) (*response.ViewDashboard, error) {
	response := new(response.ViewDashboard)
	query := `SELECT
			SUM(COALESCE(ts.total_os,0)) as target_os,
			SUM(d.os_total) as pencapaian_os,
			SUM(COALESCE(((ts.total_os * 116/100)-d.os_total)/(12-MONTH(CURDATE())+1),0)) as minimal_delta_nilai_4,
			SUM((case when u.next_pmt_date >= u.Periode then u.Plafond / CAST(TRIM(BOTH 'M' FROM u.jangka_waktu) as int) else 0 end)) as angsuran_pokok_belum_masuk,
			SUM(d.pencapaian_realisasi) as pencapaian_realisasi,
			SUM(d.sisa_suplesi) as sisa_suplesi,
			SUM(d.rincian_lunas_hutang) as rincian_lunas_hutang,
			SUM(COALESCE(d.dh_total,0)) as dh_total,
			SUM(d.dpk_total) as dpk_posisi,
			(select count(d.dpk_total) from dashboard d WHERE d.dpk_total != 0) as count_dpk_posisi,
			SUM(d.dpk_baru) as dpk_baru,
			(select count(d.dpk_baru) from dashboard d WHERE d.dpk_baru != 0) as count_dpk_baru,
			SUM(d.dpk_1) as dpk_1,
			(select count(d.dpk_1) from dashboard d WHERE d.dpk_1 != 0) as count_dpk_1,
			SUM(d.dpk_2) as dpk_2,
			(select count(d.dpk_2) from dashboard d WHERE d.dpk_2 != 0) as count_dpk_2,
			SUM(d.dpk_3) as dpk_3,
			(select count(d.dpk_3) from dashboard d WHERE d.dpk_3 != 0) as count_dpk_3,
			SUM(d.dpk_3) as dpk_calon_npl,
			(select count(d.dpk_3) from dashboard d WHERE d.dpk_3 != 0) as count_dpk_calon_npl,
			SUM(d.dpk_blm_restruk) as dpk_blm_restruk,
			(select count(d.dpk_blm_restruk) from dashboard d WHERE d.dpk_blm_restruk != 0) as count_dpk_blm_restruk,
			SUM(d.lancar_blm_jth_tempo) as lancar_blm_jatuh_tempo,
			(select count(d.lancar_blm_jth_tempo) from dashboard d WHERE d.lancar_blm_jth_tempo != 0) as count_lancar_blm_jatuh_tempo,
			SUM(d.npl_total) as npl_total,
			(select count(d.npl_total) from dashboard d WHERE d.npl_total != 0) as count_npl_total,
			SUM(d.npl_kl_kecil) as npl_Kl_baru,
			(select count(d.npl_kl_kecil) from dashboard d WHERE d.npl_kl_kecil != 0) as count_npl_kl_baru,
			SUM(d.npl_kl_total) as npl_Kl,
			(select count(d.npl_kl_total) from dashboard d WHERE d.npl_kl_total != 0) as count_npl_kl,
			SUM(d.npl_diragukan_total) as npl_diragukan_total,
			(select count(d.npl_diragukan_total) from dashboard d WHERE d.npl_diragukan_total != 0) as count_npl_diragukan_total,
			SUM(d.npl_macet) as npl_macet,
			(select count(d.npl_macet) from dashboard d WHERE d.npl_macet != 0) as count_npl_macet,
			SUM(d.npl_belum_restruk) as npl_belum_restruk,
			(select count(d.npl_belum_restruk) from dashboard d WHERE d.npl_belum_restruk != 0) as count_npl_belum_restruk
		FROM dashboard d
		join target_smk ts on ts.pn = d.Id_mantri
		join upload u on d.Id_mantri = u.pn_pengelola
		WHERE d.Id_mantri = ?;`

	err := repo.db.Get(response, query, mantri)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
}
