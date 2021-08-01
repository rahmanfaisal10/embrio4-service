package repository

import (
	"rahmanfaisal10/embrio4-service/pkg/model"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUserByUsernamePN(usernamePN string) (*model.Users, error)
	UpdateLastLogin(user model.Users) error
	CreateUser(user *model.Users) error
	UpdatePassword(user model.Users) error
	BulkUpsertCabang(request []*model.Cabang, tx *sqlx.Tx) ([]*model.Cabang, error)
	BulkUpsertUnit(request []*model.Unit, tx *sqlx.Tx) ([]*model.Unit, error)
	BUlkUpsertMantri(request []*model.Mantri, tx *sqlx.Tx) ([]*model.Mantri, error)
	BUlkUpsertMisi(request []*model.Misi, tx *sqlx.Tx) ([]*model.Misi, error)
	BUlkUpsertNasabah(request []*model.Nasabah, tx *sqlx.Tx) ([]*model.Nasabah, error)
	UploadRepository(request []*model.Upload) error
}

type repository struct {
	db *sqlx.DB
}

func InitRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
