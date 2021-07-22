package request

type RegisterRequest struct {
	UsernamePN string `json:"username_pn,omitempty"`
	Nama       string `json:"nama"`
	UnitKerja  string `json:"unit_kerja"`
	KodeBranch string `json:"kode_branch"`
	Jabatan    string `json:"jabatan"`
	Groups     string `json:"groups"`
}
