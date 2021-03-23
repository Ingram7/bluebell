package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,len=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type User struct {
	UserID   int64
	Username string
	Password string
}

func (u User) TableName() string {
	return "user"
}
