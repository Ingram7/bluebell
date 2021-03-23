package mysql

import "bluebell/models"

func CheckUserExist(name string) bool {
	user := models.User{}
	GetDB().Debug().Where("username = ?", "name").First(&user)
	if user.UserID != 0 {
		return true
	}
	return false
}

func InsertUser(user *models.User) (err error) {
	return GetDB().Debug().Create(user).Error
}
