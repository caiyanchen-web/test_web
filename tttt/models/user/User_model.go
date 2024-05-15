package user

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

//用户数据库模型

type UserInfo struct {
	Id         int `orm:"unique"`
	Name       string
	Age        int
	PhoneNum   string
	Email      string
	Gender     string
	PassWord   string
	Address    string
	Available  bool
	Picture    string    `orm:"null"`
	CreateTime time.Time `orm:"type(datetime)"`
	DeleteTime time.Time `orm:"type(datetime);size(255);null"`
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

// 获取登录表单
type UserLogin struct {
	Name     string `form:"name"`
	PassWord string `form:"pwd"`
}

// HashPassword 使用 Scrypt 算法生成密码哈希值
func HashPassWord(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

// ComparePasswords 验证密码是否匹配
func ComparePassWords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// 表单数据转换为数据库模型
func (u *UserForm) ToUserInfo(user_info *UserInfo) {
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
}
