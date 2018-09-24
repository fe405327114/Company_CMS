package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
     //db=orm.NewOrm()
	var dbhost, dbpassword, db, dbuser, dbport string
	dbhost = beego.AppConfig.String("dbhost")
	dbpassword = beego.AppConfig.String("dbpassword")
	db = beego.AppConfig.String("db")
	dbuser = beego.AppConfig.String("dbuser")
	dbport = beego.AppConfig.String("dbport")

	orm.RegisterDriver("mysql", orm.DRMySQL) //注册mysql Driver
	//构造conn连接
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	beego.Info(conn)
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)

	orm.RegisterModel(new(UserInfo),new(RoleInfo),new(ActionInfo)) //注册模型
	orm.RunSyncdb("default", false, true)
}
