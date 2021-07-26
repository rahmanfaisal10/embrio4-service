package request

type FileRequest struct {
	Periode                    string `json:"Periode"`
	Branch                     string `json:"Branch"`
	Currency                   string `json:"Currency"`
	NamaAO                     string `json:"nama AO"`
	LNType                     string `json:"LN Type"`
	NomorRekening              string `json:"Nomor Rekening"`
	NamaDebitur                string `json:"Nama Debitur"`
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
	KolektibilitasKurangLancar string `json:"KolektibilitasKurangLancar"`
	KolektibilitasDiragukan    string `json:"KOlektibilitas Diragukan"`
	KolektibilitasMacet        string `json:"Kolektibilitas Macet"`
	TunggakanPokok             string `json:"Tunggakan Pokok"`
	TunggakanBunga             string `json:"Tunggakan Bunga"`
	TunggakanPinalty           string `json:"Tunggakan Pinalty"`
	PN                         string `json:"PN"`
	NamaPN                     string `json:"Nama PN"`
	Code                       string `json:"Code"`
	Description                string `json:"Description"`
	Kol_ADK                    string `json:"Kol_ADK"`
	AvgOSHarian                string `json:"Avg OS Harian"`
	KecamatanTempatTinggal     string `json:"Kecamatan Tempat Tinggal"`
	KelurahanTempatTinggal     string `json:"Kelurahan Tempat Tinggal"`
	KodePosTempatTinggal       string `json:"Kode_pos Tempat Tinggal"`
	KecamatanTempatUsaha       string `json:"Kecamatan Tempat Usaha"`
	KelurahanTempatUsaha       string `json:"Kelurahan Tempat Usaha"`
	KodePosTempatUsaha         string `json:"Kode Pos Tempat Usaha"`
}
