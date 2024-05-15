package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"tttt/models/auth"
	database "tttt/models/database"
	"tttt/models/msg"
	"tttt/models/user"
)

type ModifyController struct {
	beego.Controller
}
type DeleteController struct {
	beego.Controller
}

func (c *ModifyController) Get() {
	//获取表单中的用户id

	var session = c.GetSession("user")
	if session == nil {
		c.Redirect("/", 302)
	} else {
		formId := c.GetString("userid")
		fmt.Println(formId)
		//转换string为int
		userId, err := strconv.Atoi(formId)
		msg.CheckErr(err)
		//通过id查询用户
		userinfo := database.SearchUserForId(userId)
		c.Data["userinfo"] = userinfo
		c.TplName = "modify.html"
	}
}

func (c *ModifyController) Post() {
	var session = c.GetSession("user")
	if session == nil {
		c.Redirect("/", 302)
	} else {
		//接收表单数据
		ID := c.GetString("id")
		id, e := strconv.Atoi(ID)
		msg.CheckErr(e)
		formUser := user.UserForm{}
		err := c.ParseForm(&formUser)
		msg.CheckErr(err)
		//对表单数据进行验证
		msg := auth.Auth(&formUser)
		if msg.Code != 0 {
			c.Data["json"] = msg.Msg
			c.ServeJSON()
		} else {
			m1 := database.UpdateUser(id, &formUser)
			if m1.Code == 0 {
				c.SetSession("user", formUser.Name)
				c.Redirect("/index", 302)
			} else {
				c.Data["json"] = m1.Msg
				c.ServeJSON()
			}
		}

	}
}

func (c *DeleteController) Get() {
	userId := c.GetString("userid")
	id, err := strconv.Atoi(userId)
	msg.CheckErr(err)
	m := database.DeleteUser(id)
	if m.Code == 0 {
		c.Redirect("/index", 302)
	} else {
		c.Data["json"] = m.Msg
		c.ServeJSON()
	}

}
