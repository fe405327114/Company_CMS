package models

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func MenuSelectdb(userId int, resp map[string]interface{}) map[string]interface{} {
	//1:获取用户的信息。
	o := orm.NewOrm()
	var userInfo = UserInfo{}
	erruI := o.QueryTable("user_info").Filter("id", userId).One(&userInfo)
	beego.Info(userId)
	if erruI != nil {
		resp["flag"] = "no"
		beego.Error(erruI)
		return resp
	}
	//获取该用户的的角色信息
	n, erruR := o.LoadRelated(&userInfo, "Roles")
	if erruR != nil {
		if n == 0 {
			beego.Info("用户现在没有角色 ")
		} else {
			resp["flag"] = "no"
			beego.Error(erruR)
			return resp
		}
	}
	//获取所有角色里的所有权限信息
	//var actionInfo=[]*models.ActionInfo{}
	var actions = []*ActionInfo{}
	for _, role := range userInfo.Roles {
		//查询到单个role里面的所有权限,多对多查询
		o.LoadRelated(role, "actions") //不能加& !!因为已经有*了
		//o.QueryTable("action_info").Filter("Roles__RoleInfo__Id",role.Id).All(&actionInfo)
		for _, action := range role.Actions {
			//将所有权限追加到切片中
			actions = append(actions, action)
		}
	}
	//对查询出的权限信息进行判断，判断一下是否为菜单权限，也就是ActionTypeEnum属性的取值是否为1.
	var menuActions = []*ActionInfo{}
	for _, action := range actions {
		if action.ActionTypeEnum == 1 {
			menuActions = append(menuActions, action)
		}
	}
	//第六：按照“用户---权限”这条线查询出用户对应的权限后，还要进行菜单权限的过滤。
	var userActions = []*UserAction{}
	m, erruA := o.QueryTable("user_action").Filter("users_id", userId).All(&userActions)
	if erruA != nil {
		if m == 0 {
			beego.Info("用户现在没有权限 ")
		} else {
			resp["flag"] = "no"
			beego.Error(erruR)
			return resp
		}
	}
	var userMenuActions = []*ActionInfo{}
	for _, userAction := range userActions {
		var actionInfoMenu=ActionInfo{}
		o.QueryTable("action_info").Filter("id",userAction.Actions.Id).Filter("" +
			"action_type_enum",1).One(&actionInfoMenu)
		//if userAction.Actions.ActionTypeEnum == 1 {
		//	userMenuActions = append(userMenuActions, userAction.Actions)
		//}
		userMenuActions=append(userMenuActions,&actionInfoMenu)
	}
	//第七：将两条线查询出的权限进行合并。
	menuActions = append(menuActions, userMenuActions...)
	//去重
	resultAction := RemoveRerepeat(menuActions)
	//第九：过滤掉登录用户禁用的权限。也就是判断登录用户中关于isPass的取值。
	//查询所有被禁止的权限和查询到的权限进行对比
	var banActions = []*UserAction{}
	x, erruB := o.QueryTable("user_action").Filter("ispass", 0).Filter("users_id", userId).All(&banActions)
	if erruB != nil {
		if x == 0 {
			beego.Info("用户现在没有被禁用的权限 ")
			resp["menus"] = resultAction
		} else {
			resp["flag"] = "no"
			beego.Error(erruB)
			return resp
		}
	}
	for i := 0; i < len(resultAction); i++ {
		for j := 0; j < len(banActions); j++ {
			if resultAction[i].Id == banActions[j].Actions.Id {
				//从切片中删除元素
				resultAction = append(resultAction[:i], resultAction[i+1:]...)
			}
		}
	}
	resp["flag"] = "ok"
	resp["menus"] = resultAction
	return resp
}
func RemoveRerepeat(menuActions []*ActionInfo) []*ActionInfo {
	var resultAction = []*ActionInfo{}
	count := 0
	for _, meAction := range menuActions {
		for _, reAction := range resultAction {
			if meAction.Id == reAction.Id {
				count++
				break
			}
		}
		if count == 0 {
			resultAction = append(resultAction, meAction)
		}
	}
	return resultAction
}
func CommonSelectdb(url, method string, c *context.Context) {
	//1：从session中获取登录用户的名称。
	//2：获取用户请求的地址与请求的方式。
	//3：根据获取的请求的地址和请求的方式，从权限表中查询出对应的权限。
	//4：如果没有找到对应的权限跳转到登录页面。
	//5：如果找到了，就要判断当前登录用户是否具有该权限，可以现按照“用户--权限”这条线进行过滤。
	//6：如果按照“用户--权限”这条线，发现登录用户具有该权限，那么就要判断是否是“禁用”还是允许。
	//7：如果是“允许”就结束整个判断，如果是“禁止”，就要按照“用户--角色--权限”这条线进行判断。
	//8：如果按照“用户--角色--权限”这条线进行过滤，用户有权限，进行访问，如果没有权限跳转到登录页面。
	//查看用户是否登录
	userId := c.Input.Session("userId")
	userName := c.Input.Session("userName")
	if userName=="" {
		c.Redirect(302, "/Login")
	}
	if userName == "Admin" {
		return
	}
	o := orm.NewOrm()
	//查询到该条权限
	var actionInfo = ActionInfo{}
	erruA := o.QueryTable("action_info").Filter("Url", url).Filter("HttpMethod", method).One(&actionInfo)
	if erruA != nil {
		beego.Error(erruA)
	}
	if actionInfo.Id < 0 {
		c.Redirect(302, "/Login")
	} else {
		//判断用户是否拥有该权限、
		var userAction = UserAction{}
		erroU := o.QueryTable("user_action").Filter("users_id", userId).Filter("actions_id"+
			"", actionInfo).One(&userAction)
		if erroU != nil {
			beego.Error(erroU)
		}
		if userAction.Id > 0 { //此处分三种情况，允许，禁止，和没权限,分别进行处理
			//用户有这个权限，判断是否被禁止
			if userAction.Ispass == 1 {
				return
			} else {
				c.Redirect(302, "/Login")
			}
		} else {
			//用户没有这条权限,再用用户--角色--权限判断
			var userInfo = UserInfo{}
			_, erruR := o.LoadRelated(&userInfo, "Roles")
			if erruR != nil {
				beego.Error(erruR)
				c.Redirect(302, "/Login")
			}
			//找到该用户所有权限
			for _, role := range userInfo.Roles {
				_, errrA := o.LoadRelated(&role, "Actions")
				if errrA != nil {
					beego.Error(errrA)
					c.Redirect(302, "/Login")
				}
				//判断是否具有该权限
				count := 0
				for _, action := range role.Actions {
					if action.Id == actionInfo.Id {
						count++
						return
					}
					if count==0{
						c.Redirect(302, "/Login")
					}
				}

			}

		}

	}
}
