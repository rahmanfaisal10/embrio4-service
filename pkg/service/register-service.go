package service

import (
	"errors"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"rahmanfaisal10/embrio4-service/pkg/request"
	"rahmanfaisal10/embrio4-service/pkg/util"

	"golang.org/x/crypto/bcrypt"
)

func (s service) RegisterService(request request.RegisterRequest) error {
	//checking data is exist on database
	user, _ := s.r.GetUserByUsernamePN(request.UsernamePN)
	if user != nil {
		err := errors.New("username_pn does exists, please using another username_pn")
		return err
	}

	//generate from request password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	newUser := &model.Users{
		UserID:     util.GeneratorString(6),
		UsernamePN: request.UsernamePN,
		Password:   string(password),
		Nama:       request.Nama,
		UnitKerja:  request.UnitKerja,
		KodeBranch: request.KodeBranch,
		Jabatan:    request.Jabatan,
	}

	err = s.r.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}
