package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"tttt/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/register", &controllers.Register{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/logout", &controllers.LogoutController{}, "*:Logout") //退出
	beego.Router("/modify", &controllers.ModifyController{})
}
