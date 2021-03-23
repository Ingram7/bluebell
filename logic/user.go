package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"

	"github.com/pkg/errors"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if mysql.CheckUserExist(p.Username) {
		return errors.New("用户已经存在")
	}
	// 生成UID
	userID := snowflake.GenID()
	// 密码加密
	u := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	mysql.InsertUser(&u)

	str := "hello world!"
	fmt.Printf("%d\n", str)

	return nil
}
