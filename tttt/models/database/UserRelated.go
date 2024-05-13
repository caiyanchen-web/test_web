package database

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"tttt/models/msg"
	"tttt/models/user"
)

func RegisterUser(u *user.UserForm) (m msg.Msg) {
	//defer msg.RecoverPanic()
	o1 := orm.NewOrm()
	userInfo := &user.User_Info{}
	u.ToUserInfo(userInfo)
	userInfo.PassWord = user.HashPassWord(u.PassWord)
	id, err := o1.Insert(userInfo)
	if err != nil {
		fmt.Println(id)
	}
	if err != nil {
		msg.CheckErr(err)
		m = msg.Msg{
			Code: 501,
			Msg:  "注册失败",
		}
	} else {
		m = msg.Msg{
			Code: 0,
			Msg:  "注册成功",
			Data: nil,
		}
	}
	return m
}
func LoginUser(u *user.UserLogin) (m msg.Msg) {
	//defer msg.RecoverPanic()
	o1 := orm.NewOrm()
	userInfo := &user.User_Info{}
	userInfo.Name = u.Name
	err := o1.Read(userInfo, "Name")
	if err != nil {
		fmt.Println(err)
		m = msg.Msg{
			Code: 601,
			Msg:  "用户名不存在",
		}
	} else if user.ComparePassWords(userInfo.PassWord, u.PassWord) {
		fmt.Println(userInfo.PassWord)
		fmt.Println(u.PassWord)
		fmt.Println(err)
		m = msg.Msg{
			Code: 602,
			Msg:  "密码错误",
		}
	} else {
		m = msg.Msg{
			Code: 0,
			Msg:  "登录成功",
		}
	}
	return m
}
