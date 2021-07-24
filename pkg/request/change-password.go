package request

type ChangePasswordRequest struct {
	UsernamePN  string `json:"username_pn,omitempty"`
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}
