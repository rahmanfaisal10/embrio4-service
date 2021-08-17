package response

type MiniDashboardResponse struct {
	DpkTotal float64 `json:"dpk_total" db:"dpk_total"`
	Status   string  `json:"status" db:"status"`
	Count    int64   `json:"count" db:"count"`
}
