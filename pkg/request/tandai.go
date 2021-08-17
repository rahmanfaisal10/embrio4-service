package request

type LogTandaiRequest struct {
	PnPengguna            string `json:"pn_pengguna" db:"pn_pengguna"`
	NamaPenguna           string `json:"nama_pengguna" db:"nama_pengguna"`
	Periode               string `json:"periode" db:"periode"`
	CifNo                 string `json:"cif_no" db:"cif_no"`
	NomorRekeningSimpanan string `json:"nomor_rekening_simpanan" db:"nomor_rekening_simpanan"`
	NomorRekeningPinjaman string `json:"nomor_rekening_pinjaman" db:"nomor_rekening_pinjaman"`
	Status                string `json:"status" db:"status"`
	TglJanjiSetor         string `json:"tgl_janji_setor" db:"tgl_janji_setor"`
}
