package service

import (
	"go-starter/src/dao"
	"go-starter/src/helper"
	"go-starter/src/model"
	"time"
)

// newly build User Information
func CreateUser(user *model.User) (err error) {
	user.Password, err = helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	if err := dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	dao.SqlSession.Model(user).Update("CreatedAt", time.Now())
	return
}

// Obtain user List
func GetAlluser() (userList []*model.User, err error) {
	if err := dao.SqlSession.Order("id asc").Find(&userList).Error; err != nil {
		return nil, err
	}
	return

}

// Delete user
func DeleteUserById(id string) (err error) {
	if err := dao.SqlSession.Where("id=?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return
}

// Find user By Id
func GetUserById(id string) (user *model.User, err error) {
	// dao.SqlSession.Model(&model.User{}).Where("id=?", id).First(&user)

	// dao.SqlSession.Raw("Select * from user where id = ?", id).Scan(&user)

	if err := dao.SqlSession.Where("id=?", id).First(&model.User{}).Error; err != nil {
		return nil, err
	}
	return
}

// Find user By Username
func GetUserByUsername(username string) (*model.User, error) {
	var usr = &model.User{}
	if err := dao.SqlSession.Where("username=?", username).First(usr).Error; err != nil {
		return nil, err
	}
	return usr, nil
}

// Update user
func UpdateUser(user *model.User) (err error) {
	if err := dao.SqlSession.Save(user).Error; err != nil {
		return err
	}
	return
}
