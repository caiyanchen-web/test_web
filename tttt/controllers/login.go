package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	database "tttt/models/database"
	"tttt/models/user"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}
func (c *LoginController) Post() {
	//定义登录结构体，获取表单数据
	u := &user.UserLogin{}
	c.ParseForm(u)
	fmt.Println(u)
	//进行密码与用户名验证
	m := database.LoginUser(u)
	if m.Code != 0 {
		c.Data["json"] = m.Msg
		c.ServeJSON()
	} else {
		c.SetSession("user", u.Name)
		c.Redirect("/index.html", 302)
	}
}
