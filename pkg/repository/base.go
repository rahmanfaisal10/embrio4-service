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
	// BulkUpsertCabang(request []*model.Cabang) ([]*model.Cabang, error)
	// BUlkUpsertMantri(request []*model.Mantri) ([]*model.Mantri, error)
}

type repository struct {
	db *sqlx.DB
}

func InitRepository(db *sqlx.DB) Repository {
	return &repository{db}
}
