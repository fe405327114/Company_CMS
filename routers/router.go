package routers

import (
	"Company_Cms/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"Company_Cms/models"
)

func init() {
	//-----------------用户信息-------------------------
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Admin/UserInfo", &controllers.UserInfoController{}, "get:Index")
	beego.Router("/Admin/UserInfo/AddUser", &controllers.UserInfoController{}, "post:AddUser")
	beego.Router("/Admin/UserInfo/GetUserInfo", &controllers.UserInfoController{}, "post:GetUserInfo")
	beego.Router("/Admin/UserInfo/DeleteUser", &controllers.UserInfoController{}, "post:DeleteUser")
	beego.Router("/Admin/UserInfo/ShowEditUser", &controllers.UserInfoController{}, "post:ShowEditUser")
	beego.Router("/Admin/UserInfo/EditUser", &controllers.UserInfoController{}, "post:EditUser")
	beego.Router("/Admin/UserInfo/AddUserRole", &controllers.UserInfoController{}, "get:AddUserRole")
	beego.Router("/Admin/UserInfo/EditRole", &controllers.UserInfoController{}, "post:EditRole")
	beego.Router("/Admin/UserInfo/ShowUserAction",&controllers.UserInfoController{},"get:ShowUserAction")
	beego.Router("/Admin/UserInfo/SetUserAction",&controllers.UserInfoController{},"post:SetUserAction")
	beego.Router("/Admin/UserInfo/DeleteUserAction",&controllers.UserInfoController{},"post:DeleteUserAction")

	//------------------角色信息--------------------------
	beego.Router("/Admin/RoleInfo", &controllers.RoleInfoController{}, "get:Index")
	beego.Router("/Admin/RoleInfo/GetRoleInfo", &controllers.RoleInfoController{}, "post:GetRoleInfo")
	beego.Router("/Admin/RoleInfo/ShowAddRole", &controllers.RoleInfoController{}, "get:ShowAddRole")
	beego.Router("/Admin/RoleInfo/AddRole", &controllers.RoleInfoController{}, "post:AddRole")
	beego.Router("/Admin/RoleInfo/DeleteRole", &controllers.RoleInfoController{}, "post:DeleteRole")
	beego.Router("/Admin/RoleInfo/ShowEditRole", &controllers.RoleInfoController{}, "get:ShowEditRole")
	beego.Router("/Admin/RoleInfo/ShowSetRoleAction",&controllers.RoleInfoController{},"get:ShowSetRoleAction")
	beego.Router("/Admin/RoleInfo/SetRoleAction" ,&controllers.RoleInfoController{},"post:SetRoleAction")

	//------------------权限信息--------------------------
	beego.Router("/Admin/ActionInfo",&controllers.ActionInfoController{},"get:Index")
	beego.Router("/Admin/ActionInfo/GetActionInfo",&controllers.ActionInfoController{},"post:GetActionInfo")
	beego.Router("/Admin/ActionInfo/ShowAddAction",&controllers.ActionInfoController{},"get:ShowAddAction")
	beego.Router("/Admin/ActionInfo/FileUp",&controllers.ActionInfoController{},"post:FileUp")
    beego.Router("/Admin/ActionInfo/AddAction",&controllers.ActionInfoController{},"post:AddAction")
	beego.Router("/Admin/ActionInfo/DeleteAction",&controllers.ActionInfoController{},"post:DeleteAction")
	beego.Router("/Admin/ActionInfo/ShowEditAction",&controllers.ActionInfoController{},"get:ShowEditAction")
	beego.Router("/Admin/ActionInfo/EditAction",&controllers.ActionInfoController{},"post:EditAction")

	//------------------页面布局--------------------------
	beego.Router("/Admin/Home/Index",&controllers.HomeController{},"get:Index")
	beego.Router("/Admin/Home/WindowIndex",&controllers.HomeController{},"get:WindowIndex")
	beego.Router("/Admin/Home/MenuSelect",&controllers.HomeController{},"post:MenuSelect")

	//----------------------------------新闻类别---------------------------------------
	beego.Router("/Admin/ArticelClass/Index",&controllers.ArticelClassController{},"get:Index")
	beego.Router("/Admin/ArticelClass/ShowAddParent",&controllers.ArticelClassController{},"get:ShowAddParent")
	beego.Router("/Admin/ArticelClass/AddParentClass",&controllers.ArticelClassController{},"post:AddParentClass")
	beego.Router("/Admin/ArticelClass/ShowArticelClass",&controllers.ArticelClassController{},"post:ShowArticelClass")
	beego.Router("/Admin/ArticelClass/ShowAddChildClass",&controllers.ArticelClassController{},"get:ShowAddChildClass")
	beego.Router("/Admin/ArticelClass/AddChildClass",&controllers.ArticelClassController{},"post:AddChildClass")
	beego.Router("/Admin/ArticelClass/ShowChildClass",&controllers.ArticelClassController{},"post:ShowChildClass")

	//------------------登录页面--------------------------
	beego.Router("/Login",&controllers.LoginController{},"get:Index")
    beego.Router("/Login/UserLogin",&controllers.LoginController{},"post:UserLogin")
	beego.InsertFilter("/Admin/*",beego.BeforeExec,CommonSelect)
}
//定义过滤路由的函数，控制请求权限
func CommonSelect(ctx *context.Context){
	url:=ctx.Request.URL.Path
	method:=ctx.Request.Method
	//开始过滤
	models.CommonSelectdb(url,method,ctx)
}
