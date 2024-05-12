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

func (c *Register) Get() {
	c.TplName = "register.html"

}

// 用户注册
func (c *Register) Post() {
	//上传头像
	f, h, e1 := c.GetFile("image")
	defer f.Close()
	//头像存放路径
	img_path := "static/img/"
	//拼接头像存储路径
	imagename := path.Join(img_path, h.Filename)
	//存储头像
	c.SaveToFile("image", imagename)

	msg.CheckErr(e1)

	defer msg.RecoverPanic()
	//解析页面表单
	user := modelsuser.UserForm{}
	user.Img = imagename
	if err := c.ParseForm(&user); err != nil {
		msg.CheckErr(err)
	}

	fmt.Println(user)

	m := auth.Auth(&user)

	if m.Code == 0 {
		m1 := database.RegisterUser(&user)
		fmt.Println(m1)
		if m1.Code == 0 {
			//注册成功后跳转至主页
			c.Redirect("/", 302)
		}
	} else {
		fmt.Println(m)
	}

}
