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
	fmt.Println(userInfo)
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
