package response

import "time"

type MiniDashboardResponse struct {
	Periode  time.Time `json:"periode" db:"periode"`
	DpkTotal float64   `json:"dpk_total" db:"dpk_total"`
	Status   string    `json:"status" db:"status"`
	Count    int64     `json:"count" db:"count"`
}
