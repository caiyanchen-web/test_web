package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"path"
	"tttt/models/auth"
	database "tttt/models/database"
	"tttt/models/msg"
	modelsuser "tttt/models/user"
)

type Register struct {
	beego.Controller
}

// 用户注册
func (c *Register) RegisterUser() {
	c.TplName = "register.html"
	defer msg.RecoverPanic()
	//解析页面表单
	user := modelsuser.UserForm{}
	if err := c.ParseForm(&user); err != nil {
		msg.CheckErr(err)
	}
	//头像存放路径
	img_path := "static/img/"
	//上传头像
	f, h, e1 := c.GetFile("image")
	msg.CheckErr(e1)
	defer f.Close()
	//拼接头像存储路径
	imagename := path.Join(img_path, h.Filename)
	//存储头像
	c.SaveToFile("image", imagename)
	user.Img = imagename
	fmt.Println(user)

	m := auth.Auth(&user)
	m1 := database.RegisterUser(&user)
	if m.Code == 0 {
		if m1.Code == 0 {
			//注册成功后跳转至主页
			c.Redirect("/", 302)
		}
	} else {
		fmt.Println(m)
		fmt.Println(m1)
	}

}
