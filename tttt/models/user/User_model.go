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

// 表单数据转换为数据库模型
func (u *UserForm) ToUserInfo() (user_info *User_Info) {
	user_info.Name = u.Name
	user_info.Age = u.Age
	user_info.PhoneNum = u.PhoneNum
	user_info.Email = u.Email
	user_info.Gender = u.Gender
	user_info.PassWord = u.PassWord
	user_info.Address = u.Address
	user_info.Available = true
	user_info.Picture = u.Img
	user_info.CreateTime = time.Now()
	user_info.DeleteTime = time.Time{}
	return user_info
}
