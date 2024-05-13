package controllers

import beego "github.com/beego/beego/v2/server/web"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	loginsession := c.GetSession("user")
	if loginsession != nil {
		c.Data["image"] = "static/img/1.png"
		c.TplName = "Include/footer.html"
		c.TplName = "index.html"
	} else {
		c.Redirect("/", 302)
	}
}

func (c *IndexController) logout() {
	c.DelSession("user")
	c.Redirect("/", 302)
}
