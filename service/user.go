package service

import (
	"errors"
	"gorm.io/gorm"
	"xiaowei_list/global"
	"xiaowei_list/model/database"
)

// Helper

// 数据库操作

// CreateUser 创建用户
func CreateUser(user *database.User) (err error) {
	if err = global.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID 根据用户 ID 查询某个用户
func GetUserByID(ID int) (user database.User, notFound bool) {
	err := global.DB.First(&user, ID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else {
		return user, false
	}
}

// GetUserByUsername 根据用户名查询某个用户
func GetUserByUsername(username string) (user database.User, notFound bool) {
	err := global.DB.Where("Username = ?", username).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else {
		return user, false
	}
}

// SaveUser 根据信息保存用户
func SaveUser(user *database.User) (err error) {
	err = global.DB.Save(user).Error
	return err
}

// DeleteUserByID 根据uid 删除用户
func DeleteUserByID(ID int) (err error) {
	if err = global.DB.Where("UserId = ?", ID).Delete(database.User{}).Error; err != nil {
		return err
	}
	return nil
}
