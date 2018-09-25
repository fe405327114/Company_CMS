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
	beego.AddFuncMap("CheckUserAction",CheckUserAction)
	beego.AddFuncMap("CheckUserActionIspass",CheckUserActionIspass)
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
//判断用户是否具有某个权限
func CheckUserAction(userExtActionList[]*models.UserAction,actionId int)(b bool) {
	b=false
	for i:=0;i<len(userExtActionList);i++ {
		if userExtActionList[i].Actions.Id==actionId {
			b=true
			break
		}
	}
	return
}
//判断用户具有某个权限是禁止还是允许
func CheckUserActionIspass(userExtActionList[]*models.UserAction)(b bool) {
	b = false
	for i := 0; i < len(userExtActionList); i++ {
			if userExtActionList[i].Ispass == 1 {
				b = true
			}
			break //注意break的位置
	}
	return
}