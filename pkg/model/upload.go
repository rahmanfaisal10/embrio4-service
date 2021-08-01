package model

type Upload struct {
	ID                         int    `json:"id"`
	Periode                    string `json:"Periode" db:"periode"`
	Region                     string `json:"region" db:"region"`
	Mainbranch                 string `json:"Main Branch" db:"Main Branch"`
	Branch                     string `json:"Branch" db:"Branch"`
	Currency                   string `json:"Currency" db:"Currency"`
	NamaAO                     string `json:"Nama AO" db:"Nama Ao"`
	LNType                     string `json:"LN Type" db:"LN Type"`
	NomorRekening              string `json:"Nomor rekening" db:"Nomor rekening"`
	NamaDebitur                string `json:"Nama Debitur" db:"Nama Debitur"`
	AlamatIdentitas            string `json:"Alamat Identitias" db:"Alamat Identitas"`
	KodePosIdentitas           string `json:"Kode Pos Identitas" db:"Kode Pos Identitas"`
	AlamatKantor               string `json:"Alamat Kantor" db:"Alamat Kantor"`
	KodePosKantor              string `json:"Kode Pos Kantor" db:"Kode Pos Alamat"`
	Plafond                    string `json:"Plafond" db:"Plafond"`
	NextPmtDate                string `json:"Next Pmt Date" db:"Next Pmt Date"`
	NextIntPmtDate             string `json:"Next Int Pmt Date" db:"Next Int Pmt Date"`
	Rate                       string `json:"Rate" db:"Rate"`
	TglMenunggak               string `json:"Tgl Menunggak" db:"Tgl Menunggak"`
	TglRealisasi               string `json:"Tgl Realisasi" db:"Tgl Realisasi"`
	TglJatuhTempo              string `json:"Tgl Jatuh Tempo" db:"Tgl Jatuh Tempo"`
	JangkaWaktu                string `json:"Jangka Waktu" db:"Jangka Waktu"`
	FlagRestruk                string `json:"Flag Restruk" db:"Flag Restruk"`
	CIFNO                      string `json:"CIFNO" db:"CIFNO"`
	KolektibilitasLancar       string `json:"Kolektibilitas Lancar" db:"Kolektibilitas Lancar"`
	KolektibilitasDPK          string `json:"Kolektibilitas DPK" db:"Kolektibilitas DPK"`
	KolektibilitasKurangLancar string `json:"Kolektibilitas Kurang Lancar" db:"Kolektibilitas Kurang Lancar"`
	KolektibilitasDiragukan    string `json:"KOlektibilitas Diragukan" db:"Kolektibilitas Diragukan"`
	KolektibilitasMacet        string `json:"Kolektibilitas Macet" db:"Kolektibilitas Macet"`
	TunggakanPokok             string `json:"Tunggakan Pokok" db:"Tunggakan Pokok"`
	TunggakanBunga             string `json:"Tunggakan Bunga" db:"Tunggakan Bunga"`
	TunggakanPinalty           string `json:"Tunggakan Pinalty" db:"Tunggakan Pinalty"`
	PNPengelola                string `json:"PN Pengelola" db:"PN Pengelola"`
	NamaPengelola              string `json:"NAMA PENGELOLA" db:"NAMA PENGELOLA"`
}
