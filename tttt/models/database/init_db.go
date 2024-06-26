package database

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"tttt/models/msg"
	"tttt/models/user"
)

func init() {
	//defer msg.RecoverPanic()
	//加载配置文件中的数据库配置
	dbuser, user_err := beego.AppConfig.String("dbuser")
	msg.CheckErr(user_err)

	dbpassword := "Yjk#2u7*W5" //数据库密码
	//数据库路径
	db, db_err := beego.AppConfig.String("dbpath")
	msg.CheckErr(db_err)
	//数据库地址拼接
	database := fmt.Sprint(dbuser + ":" + dbpassword + db)
	fmt.Println(database)
	//数据库注册
	online := orm.RegisterDriver("mysql", orm.DRMySQL)
	if online != nil {
		log.Panic("数据库连接失败", online)
	}
	orm.RegisterDataBase("default", "mysql", database)
	//创建数据库
	createDataBase()
}
func tableName() string { return "userinfo" }

func createDataBase() {
	tableName()
	orm.RegisterModel(new(user.UserInfo))
	orm.RunSyncdb("default", false, true)
}
