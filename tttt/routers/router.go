package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"tttt/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/register", &controllers.Register{})
}
