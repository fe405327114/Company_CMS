package routers

import (
	"ItcastCms/controllers"
	"github.com/astaxie/beego"
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

}
