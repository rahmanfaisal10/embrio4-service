package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/response"

	"github.com/labstack/gommon/log"
)

func (r *repository) GetUserByUsernamePN(usernamePN string) (*model.Users, error) {
	user := new(model.Users)
	query := `SELECT * FROM users WHERE username_pn = ?`
	err := r.db.Get(user, query, usernamePN)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) UpdateLastLogin(user model.Users) error {
	queryUpdate := `UPDATE users SET last_login = NOW() WHERE username_pn=?`
	_, err := r.db.Exec(queryUpdate, user.UsernamePN)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (r *repository) CreateUser(user *model.Users) error {
	query := `INSERT INTO users(username_pn, password, nama, unit_kerja, kode_branch, jabatan, created_at, created_by, updated_at, updated_by, last_login, groups, phone_number)
			VALUES ( :username_pn, :password, :nama, :unit_kerja, :kode_branch, :jabatan, NOW(), 'admin', NOW(), 'admin', :last_login, :groups, :phone_number);`

	_, err := r.db.NamedExec(query, user)
	if err != nil {
		err = errors.New("can't to Creates data to Table user with : " + err.Error())
		return err
	}
	return nil
}

func (r *repository) UpdatePassword(user model.Users) error {
	queryUpdate := `UPDATE users SET password = ?, updated_at = NOW() WHERE username_pn= ?`
	_, err := r.db.Exec(queryUpdate, user.Password, user.UsernamePN)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (r *repository) GetAllMantriFromUpload() ([]response.RegisterHelperResponse, error) {
	query := `SELECT DISTINCT u.pn_pengelola, u.nama_pengelola, u.branch FROM upload u WHERE u.pn_pengelola != ''`
	pnPengelola := make([]response.RegisterHelperResponse, 0)

	err := r.db.Select(&pnPengelola, query)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fmt.Println(pnPengelola)
	return pnPengelola, nil
}
