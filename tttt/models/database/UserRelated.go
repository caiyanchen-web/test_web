package database

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"tttt/models/msg"
	"tttt/models/user"
)

// 注册用户
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

//展示所有用户

func ListUser() []user.User_Info {
	var userinfo []user.User_Info
	qs := orm.NewOrm().QueryTable(new(user.User_Info))
	_, err := qs.All(&userinfo)
	msg.CheckErr(err)
	return userinfo
}

// 通过名称查询用户
func SearchUserForName(name string) *user.User_Info {
	o1 := orm.NewOrm()
	u1 := user.User_Info{Name: name}
	err := o1.Read(&u1, "name")
	msg.CheckErr(err)
	return &u1
}

// 通过id查询用户
func SearchUserForId(id int) *user.User_Info {
	o1 := orm.NewOrm()
	u1 := user.User_Info{Id: id}
	err := o1.Read(&u1, "id")
	msg.CheckErr(err)
	return &u1
}

// 更新用户信息
func UpdateUser(u *user.User_Info) {
	o1 := orm.NewOrm()

	_, err := o1.Update(u)
	msg.CheckErr(err)
}
