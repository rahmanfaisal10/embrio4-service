package repository

import (
	"database/sql"
	"rahmanfaisal10/embrio4-service/pkg/model"

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
}

type repository struct {
	db *sqlx.DB
}

func InitRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
