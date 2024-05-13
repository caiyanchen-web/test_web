package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "tttt/models/database"
	_ "tttt/routers"
)

func main() {
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//连接数据库以及初始化表
	beego.Run()

}
