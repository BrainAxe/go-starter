package service

import (
	"errors"
	"go-starter/src/dao"
	"go-starter/src/model"

	"go-starter/src/helper"
)

// Authenticate User
func AuthenticateUser(username string, password string) (*model.User, error) {
	var usr = &model.User{}
	if err := dao.SqlSession.Where("username=?", username).First(usr).Error; err != nil {
		return nil, errors.New("Username doesn't exists!")
	}

	if helper.CheckPasswordHash(password, usr.Password) {
		return usr, nil
	} else {
		return nil, errors.New("Password doesn't match!")
	}
}
