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
	userInfo := &user.UserInfo{}
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
	userInfo := &user.UserInfo{}
	userInfo.Name = u.Name
	fmt.Println(user.HashPassWord(u.PassWord))
	err := o1.Read(userInfo, "name")
	fmt.Println(userInfo.PassWord)
	if err != nil {
		m = msg.Msg{
			Code: 601,
			Msg:  "用户名不存在",
		}
	} else if !user.ComparePassWords(userInfo.PassWord, u.PassWord) {

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

func ListUser() []user.UserInfo {
	var userinfo []user.UserInfo
	qs := orm.NewOrm().QueryTable(new(user.UserInfo))
	_, err := qs.All(&userinfo)
	msg.CheckErr(err)
	return userinfo
}

// 通过名称查询用户
func SearchUserForName(name string) *user.UserInfo {
	o1 := orm.NewOrm()
	u1 := user.UserInfo{Name: name}
	err := o1.Read(&u1, "name")
	msg.CheckErr(err)
	return &u1
}

// 通过id查询用户
func SearchUserForId(id int) *user.UserInfo {
	o1 := orm.NewOrm()
	u1 := user.UserInfo{Id: id}
	err := o1.Read(&u1, "id")
	msg.CheckErr(err)
	return &u1
}

// 更新用户信息
func UpdateUser(id int, u *user.UserForm) (m msg.Msg) {
	defer msg.RecoverPanic()
	o1 := orm.NewOrm()
	userInfo := &user.UserInfo{Id: id}
	u.ToUserInfo(userInfo)
	userInfo.PassWord = user.HashPassWord(u.PassWord)
	fmt.Println(userInfo.PassWord)
	_, err := o1.Update(userInfo)
	msg.CheckErr(err)
	if err != nil {
		msg.CheckErr(err)
		m = msg.Msg{
			Code: 502,
			Msg:  "更新失败",
		}
	} else {
		m = msg.Msg{
			Code: 0,
			Msg:  "更新成功",
			Data: nil,
		}
	}
	return m
}

//删除用户

func DeleteUser(id int) (m msg.Msg) {
	defer msg.RecoverPanic()
	o1 := orm.NewOrm()
	userInfo := &user.UserInfo{Id: id}
	_, err := o1.Delete(userInfo)
	msg.CheckErr(err)
	if err != nil {
		msg.CheckErr(err)
		m = msg.Msg{
			Code: 503,
			Msg:  "删除失败",
		}
	} else {
		m = msg.Msg{
			Code: 0,
			Msg:  "删除成功",
			Data: nil,
		}
	}
	return m
}
