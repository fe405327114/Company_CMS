package main

import (
	_ "ItcastCms/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"ItcastCms/models"
)

func main() {
	beego.AddFuncMap("CheckId",CheckId)
	beego.AddFuncMap("CheckRoleAction",CheckRoleAction)
	beego.Run()
}

func  CheckRoleAction(hasActionList []*models.ActionInfo,actionId int)(b bool)  {
	b=false
	for i:=0;i<len(hasActionList) ;i++  {
		if hasActionList[i].Id==actionId {
			b=true
			break
		}
	}
	return
}
//创建一个模板函数，映射
//将用户已经有的权限id和全部的id进行对比，一样的返回true
func CheckId(hasRoles []*models.RoleInfo, roleId int)(b bool) {
	       b=false
           for i:=0;i<len(hasRoles);i++{
           	if hasRoles[i].Id==roleId{
           		b=true
           		return
			}
		   }
		   return
}
