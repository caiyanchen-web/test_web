package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"tttt/models/database"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	loginsession := c.GetSession("user")
	searchname := c.GetString("searchname")
	if loginsession != nil && searchname == "" {
		c.Data["image"] = "static/img/1.png"
		c.TplName = "Include/footer.html"
		users := database.ListUser()
		//u := database.SearchUser("蔡延辰")
		c.Data["Data"] = &users
		c.TplName = "index.html"
	} else if searchname != "" {
		fmt.Println(1111)
		u := database.SearchUser("蔡延辰")
		c.Data["Id"] = &u.Id
		c.Data["Name"] = &u.Name
		c.Data["Age"] = &u.Age

		c.TplName = "index.html"

	} else {
		c.Redirect("/", 302)
	}
}
