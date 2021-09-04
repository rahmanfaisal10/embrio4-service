package response

type RegisterHelperResponse struct {
	PnPengelola   string `json:"pn_pengelola" db:"pn_pengelola"`
	NamaPengelola string `json:"nama_pengelola" db:"nama_pengelola"`
	Branch        string `json:"branch" db:"branch"`
}
