package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"tttt/models/database"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	loginsession := c.GetSession("user")
	searchname := c.GetString("searchname")
	//如果前端查询返回的姓名为空并且session不为空则返回所有用户
	if loginsession != nil && searchname == "" {
		c.Data["image"] = "static/img/1.png"
		c.TplName = "Include/footer.html"
		users := database.ListUser()
		c.Data["Data"] = &users
		c.TplName = "index.html"
		//如果前端查询返回的姓名不为空则返回查询的用户并渲染到模板
	} else if loginsession != nil && searchname != "" {
		u := database.SearchUserForName(searchname)
		c.Data["user"] = &u
		c.TplName = "index.html"

	} else {
		c.Redirect("/", 302)
	}
}
