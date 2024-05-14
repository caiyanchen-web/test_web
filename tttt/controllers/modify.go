package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"tttt/models/database"
	"tttt/models/msg"
)

type ModifyController struct {
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
