package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


//文章类别
type ArticelClass struct {
	Id int//主键
	ClassName string//类别名称
	ParentId int //父类别的编号
	CreateUserId int //创建类别的用户编号
	CreateDate time.Time //创建时间
	DelFlag int //删除标记
	Remark string //备注
}
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

	orm.RegisterModel(new(UserInfo),new(RoleInfo),new(ActionInfo),new(UserAction),new(ArticelClass)) //注册模型
	orm.RunSyncdb("default", false, true)
}
