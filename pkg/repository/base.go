package repository

import (
	"database/sql"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUserByUsernamePN(usernamePN string) (*model.Users, error)
	UpdateLastLogin(user model.Users) error
	CreateUser(user *model.Users) error
	UpdatePassword(user model.Users) error
	BulkUpsertCabang(tx *sql.Tx) error
	BulkUpsertUnit(tx *sql.Tx) error
	BUlkUpsertMantri(tx *sql.Tx) error
	BUlkUpsertMisi(request []*model.Misi, tx *sqlx.Tx) error
	BUlkUpsertNasabah(tx *sql.Tx) error
	UploadRepository(request []*model.Upload) error
	QueryTotalOS(mantri string) (totalOS response.TotalOS, err error)
	QuerySisaSuplesi(mantri string) (sisaSuplesi response.SisaSuplesiResponse, err error)
	QueryRincianLunasHutang(mantri string) (sisaLunasHutang response.SisaLunasHutangResponse, err error)
	GetAllMantri() (mantri []*model.Mantri, err error)
	PencapaianRealisasi(mantri string) (pencapaianRealisasi float64, err error)
	GetAllUpload() ([]*model.Upload, error)
	InsertDashboard() error
	UploadBulanSeblumnya(cif string) (*float64, error)
	RincianLunasHutang(cif string) (*float64, error)
	ViewDashboard(mantri string) (*response.ViewDashboard, error)
}

type repository struct {
	db *sqlx.DB
}

func InitRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
