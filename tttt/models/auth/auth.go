package auth

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"tttt/models/msg"
	u1 "tttt/models/user"
)

// 验证用户密码与确认密码是否一直
func Auth(u *u1.UserForm) (m *msg.Msg) {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	// 正则调用规则
	//手机格式验证
	PhoneregRuler := "^1[345789]{1}\\d{9}$"
	Phonereg := regexp.MustCompile(PhoneregRuler)
	//邮箱格式验证
	EmailregRuler := "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	Emailreg := regexp.MustCompile(EmailregRuler)

	//验证密码不得小于8位
	PasswrdRuler := "^.{8,}$"
	Passwordreg := regexp.MustCompile(PasswrdRuler)

	if !Phonereg.MatchString(u.PhoneNum) {
		m = msg.NewMsg(201, "手机号格式不正确", nil)
	} else if !Emailreg.MatchString(u.Email) {
		m = msg.NewMsg(301, "邮箱格式不正确", nil)
	} else if u.PassWord != u.ConfirmPassword {
		m = msg.NewMsg(101, "密码与确认密码不一致", nil)
	} else if !Passwordreg.MatchString(u.PassWord) {
		m = msg.NewMsg(401, "密码不能小于8位", nil)
	} else if IsNull(u.Address) {
		m = msg.NewMsg(501, "地址不能为空", nil)
	} else if &u.Age == nil {
		m = msg.NewMsg(502, "年龄不能为空", nil)
	} else {
		m = msg.NewMsg(0, "注册成功", nil)
	}
	return m
}

// 为空验证
func IsNull(s string) bool {
	if s == "" {
		return false
	}
	return true
}

// 密码加密
func HashPassWord(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	msg.CheckErr(err)
	return string(hashedPassword)
}
