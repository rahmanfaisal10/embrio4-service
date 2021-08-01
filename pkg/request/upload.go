package request

import "rahmanfaisal10/embrio4-service/pkg/model"

type FileRequest struct {
	Periode                    string `json:"Periode"`
	Region                     string `json:"region"`
	Mainbranch                 string `json:"Main Branch"`
	Branch                     string `json:"Branch"`
	Currency                   string `json:"Currency"`
	NamaAO                     string `json:"Nama AO"`
	LNType                     string `json:"LN Type"`
	NomorRekening              string `json:"Nomor rekening"`
	NamaDebitur                string `json:"Nama Debitur"`
	AlamatIdentitas            string `json:"Alamat Identitias"`
	KodePosIdentitas           string `json:"Kode Pos"`
	AlamatKantor               string `json:"Alamat Kantor"`
	KodePosKantor              string `json:"Kode Pos"`
	Plafond                    string `json:"Plafond"`
	NextPmtDate                string `json:"Next Pmt Date"`
	NextIntPmtDate             string `json:"Next Int Pmt Date"`
	Rate                       string `json:"Rate"`
	TglMenunggak               string `json:"Tgl Menunggak"`
	TglRealisasi               string `json:"Tgl Realisasi"`
	TglJatuhTempo              string `json:"Tgl Jatuh Tempo"`
	JangkaWaktu                string `json:"Jangka Waktu"`
	FlagRestruk                string `json:"Flag Restruk"`
	CIFNO                      string `json:"CIFNO"`
	KolektibilitasLancar       string `json:"Kolektibilitas Lancar"`
	KolektibilitasDPK          string `json:"Kolektibilitas DPK"`
	KolektibilitasKurangLancar string `json:"Kolektibilitas Kurang Lancar"`
	KolektibilitasDiragukan    string `json:"KOlektibilitas Diragukan"`
	KolektibilitasMacet        string `json:"Kolektibilitas Macet"`
	TunggakanPokok             string `json:"Tunggakan Pokok"`
	TunggakanBunga             string `json:"Tunggakan Bunga"`
	TunggakanPinalty           string `json:"Tunggakan Pinalty"`
	PNPengelola                string `json:"PN Pengelola"`
	NamaPengelola              string `json:"NAMA PENGELOLA"`
}

type UploadRepository struct {
	AksiMisi        []*model.AksiMisi
	Cabang          []*model.Cabang
	Cicilan         []*model.Cicilan
	DetailCicilan   []*model.DetailCicilan
	Detailtunggakan []*model.DetailTunggakan
	JenisAksi       []*model.JenisAksi
	JenisPinjaman   []*model.JenisPinjaman
	Kolektebilitas  []*model.Kolektebilitas
	Mantri          []*model.Mantri
	Misi            []*model.Misi
	Nasabah         []*model.Nasabah
	Pinjaman        []*model.Pinjaman
	Tunggakan       []*model.Tunggakan
	Unit            []*model.Unit
}
