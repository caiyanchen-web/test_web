package user

import (
	"time"
)

//用户数据库模型

type User_Info struct {
	Id         int `orm:"unique"`
	Name       string
	Age        int
	PhoneNum   string
	Email      string
	Gender     string
	PassWord   string
	Address    string
	Available  bool
	Picture    string
	CreateTime time.Time `orm:"type(datetime)"`
	DeleteTime time.Time `orm:"type(datetime)"`
}

// 用户表单模型
type UserForm struct {
	Name            string `form:"name"`
	Age             int    `form:"age"`
	PhoneNum        string `form:"phoneNum"`
	Email           string `form:"email"`
	Gender          string `form:"gender"`
	PassWord        string `form:"pwd"`
	ConfirmPassword string `form:"pwdok"`
	Address         string `form:"address"`
	Img             string
}
