package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"path"
	"tttt/models/msg"
	usermodels "tttt/models/user"
)

type Register struct {
	beego.Controller
}

// 用户注册
func (c *Register) RegisterUser() {
	c.TplName = "register.html"

	defer msg.RecoverPanic()

	//解析页面表单
	user := usermodels.UserForm{}
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
	//注册成功后跳转至主页
	c.Redirect("/", 302)
}
