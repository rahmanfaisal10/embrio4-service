package model

import "time"

type Cicilan struct {
	ID              int       `json:"id" db:"id"`
	IDPinjaman      int       `json:"id_pinjaman" db:"id_pinjaman"`
	Nominal         int       `json:"nominal" db:"nominal"`
	NextPaymentDate time.Time `json:"next_payment_date" db:"next_payment_date"`
	Status          int       `json:"status" db:"status"`
	Description     string    `json:"description" db:"description"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
