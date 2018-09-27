package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type UserAction struct {
	Id int
	Ispass int
	Users *UserInfo `orm:"rel(fk)"`
	Actions *ActionInfo `orm:"rel(fk)"`
}

func ShowUserActiondb(userId int,resp map[string]interface{}) map[string]interface{} {
		//查询用户信息
		o:=orm.NewOrm()
		var userInfo=UserInfo{}
		erruI:=o.QueryTable("user_info").Filter("id",userId).Filter("del_flag",0).One(&userInfo)
		if erruI!=nil{
			resp["flag"]="no"
			beego.Error(erruI)
			return resp
		}
	//	//查询用户拥有的权限信息
	//	var hasActions=[]*ActionInfo{}
	//	n,errhA:=o.QueryTable("action_info").Filter("Users__UserInfo__Id" +
	//		"",userId).Filter("del_flag",0).All(&hasActions)

     //查询到用户拥有的权限(通过中间表操作)
     var hasActions=[]*UserAction{}
     n,erruA:=o.QueryTable("user_action").Filter("users_id",userId).All(&hasActions)
     if erruA!=nil{
     	if n==0{
     		beego.Info("用户暂时没权限")
		}else {
			resp["flag"]="no"
					beego.Error(erruA)
						return resp
		}
	 }
	//查询所有的权限
		var allActions=[]*ActionInfo{}
		m,erraA:=o.QueryTable("action_info").Filter("del_flag",0).All(&allActions)
		if erraA!=nil{
			if m==0{
				beego.Info("权限表暂时没权限")
			}else {
				resp["flag"]="no"
				beego.Error(erraA)
				return resp
			}
		}
		resp["flag"]="ok"
		resp["userInfo"]=&userInfo
		resp["userExtActions"]=&hasActions
		resp["allActions"]=&allActions
		return resp
}
func SetUserActiondb(userId,actionId,Ispass int,resp map[string]interface{})map[string]interface{} {
	o := orm.NewOrm()
	//查询该用户的权限id为action的权限
	var userAction = UserAction{}
	erraI := o.QueryTable("user_action").Filter("actions_id", actionId).Filter("users_id"+
		"", userId).One(&userAction)
	if erraI != nil {
		resp["flag"] = "no"
		beego.Error(erraI)
	}
	if userAction.Id > 0 {
		//表示该用户原本就有该权限，直接进行修改
		userAction.Ispass = Ispass
		_, erraTE := o.Update(&userAction)
		if erraTE != nil {
			resp["flag"] = "no"
			beego.Error(erraTE)
			return resp
		}
	} else {
		//没有权限就进行添加
		//查询到该用户
		var userInfo= UserInfo{}
		erruI := o.QueryTable("user_info").Filter("id", userId).Filter("del_flag", 0).One(&userInfo)
		if erruI != nil {
			resp["flag"] = "no"
			beego.Error(erruI)
			return resp
		}
		//查询到该权限
		var actionInfo= ActionInfo{}
		erraI := o.QueryTable("action_info").Filter("id", actionId).Filter("del_flag", 0).One(&actionInfo)
		if erraI != nil {
			resp["flag"] = "no"
			beego.Error(erraI)
			return resp
		}
		userAction.Users = &userInfo
		userAction.Actions = &actionInfo
		userAction.Ispass = Ispass
		o.Insert(&userAction)
	}
	//获得多对多对象操作
	//		m2m:=o.QueryM2M(&userInfo,"Actions")
	//		_,errm2m:=m2m.Add(&actionInfo)
	resp["flag"]="ok"
   return resp
}
func DeleteUserActiondb(userId,actionId int,resp map[string]interface{})map[string]interface{}{
	o:=orm.NewOrm()
	var userAction=UserAction{}
	err:=o.QueryTable("user_action").Filter("users_id",userId).Filter("" +
		"actions_id",actionId).One(&userAction)
	if err!=nil{
		resp["falg"]="no"
		beego.Error(err)
		return resp
	}
	_,errDel:=o.Delete(&userAction)
	if errDel!=nil{
		resp["flag"]="no"
		beego.Error(errDel)
		return resp
	}
	resp["flag"]="ok"
	return resp
}