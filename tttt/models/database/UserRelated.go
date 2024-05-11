package database

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"tttt/models/msg"
	"tttt/models/user"
)

func RegisterUser(u *user.UserForm) (m msg.Msg) {
	defer msg.RecoverPanic()
	o1 := orm.NewOrm()
	userInfo := u.ToUserInfo()
	fmt.Println(userInfo)
	id, err := o1.Insert(userInfo)
	msg.CheckErr(err)
	if err != nil {
		m = msg.Msg{
			Code: 501,
			Msg:  "注册失败",
		}
	}
	fmt.Println(id)
	return m
}
