package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/astaxie/beego"
	"strconv"
)
//角色信息
type RoleInfo struct {
	Id int
	RoleName string `orm:"size(32)"`
	Remark string
	DelFlag int
	AddDate time.Time
	ModifDate time.Time
	Users []*UserInfo`orm:"reverse(many)"`
	Actions []*ActionInfo`orm:"rel(m2m)"`
}

func GetRoleInfodb(Roles *[]RoleInfo,pageSize,pageIndex int,resp map[string]interface{})map[string]interface{}{
    o:=orm.NewOrm()
	start:=(pageIndex-1)*pageSize
   _,err:= o.QueryTable("role_info").Filter("del_flag",0).OrderBy(
    	"id").Limit(pageSize,start).All(Roles)
   count,_:=o.QueryTable("role_info").Count()
   if err!=nil{
   	resp["flag"]="no"
   }else{
   	resp["flag"]="ok"
   	resp["rows"]=Roles
   	resp["total"]=count
   }
   return resp
   }
func AddRoledb(roleName,roleRemark string,resp map[string]interface{})map[string]interface{}{
	var roleInfo=RoleInfo{}
	roleInfo.RoleName=roleName
	roleInfo.Remark=roleRemark
	roleInfo.AddDate=time.Now()
	roleInfo.ModifDate=time.Now()
	roleInfo.DelFlag=0
	o:=orm.NewOrm()
	_,err:=o.Insert(&roleInfo)
	if err!=nil{
		resp["flag"]="no"
	}else {
		resp["flag"]="ok"
		resp["rows"]=&roleInfo
	}
	return resp

}
func DeleteRoledb(strRoleId string,resp map[string]interface{})map[string]interface{}{
	tempId:=strings.Split(strRoleId,",")
	o:=orm.NewOrm()
	//循环删除
	   var roleInfo=RoleInfo{}
	   for i:=0;i<len(tempId);i++{
		roleId,_:=strconv.Atoi(tempId[i])
		roleInfo.Id=roleId
		_,err:=o.Delete(&roleInfo)
		if err==nil{
			resp["flag"]="ok"
		}else {
			resp["flag"]="no"
			beego.Error("删除角色失败",err)
		}
	}
	return resp
}
func ShowEditRoledb(id int,roleInfo *RoleInfo,resp map[string]interface{})map[string]interface{}{
	//查询该id
	o:=orm.NewOrm()
	err:=o.QueryTable("role_info").Filter("id",id).One(roleInfo)
	if err==nil{
		resp["flag"]="ok"
	}else {
		resp["flag"]="no"
	}
	return resp
}
func EditRoledb(editName,editRemark string,id int,resp map[string]interface{})map[string]interface{}{
	//先读取出该条信息，在更新
	o:=orm.NewOrm()
	var roleInfo=RoleInfo{}
	//读取
	err:=o.QueryTable("role_info").Filter("id",id).Limit(1).One(&roleInfo)
	if err!=nil{
		resp["flag"]="no"
		return resp
	}
	//更新
	roleInfo.RoleName=editName
	roleInfo.Remark=editRemark
	_,err1:=o.Update(&roleInfo)
	if err1==nil{
		resp["flag"]="ok"
	}else{
		resp["flag"]="no"
	}
	return resp
}
func ShowSetRoleActiondb(roleId int,resp map[string]interface{})map[string]interface{} {
    var roleInfo=RoleInfo{}
    o:=orm.NewOrm()
    err:=o.QueryTable("role_info").Filter("id",roleId).One(&roleInfo)
    if err!=nil{
    	resp["flag"]="no"
    	beego.Error(err)
    	return resp
	}
	//查询角色拥有的权限
	var hasActions=[]*ActionInfo{}
	n,haserr:=o.QueryTable("action_info").Filter("Roles__RoleInfo__Id",roleId).Filter("" +
		"del_flag",0).All(&hasActions)
	if haserr!=nil{
		if n==0{
			beego.Info("用户还没有权限")
		}else{
			resp["flag"]="no"
			beego.Error(err)
			return resp
		}
	}
	//查询全部的角色
	var allActions=[]*ActionInfo{}
	m,allerr:=o.QueryTable("action_info").Filter("del_flag",0).All(&allActions)
	if allerr!=nil{
		if m==0{
			beego.Info("权限表里没有信息")
		}else {
			resp["flag"]="no"
			beego.Error(err)
			return resp
		}
	}
	resp["hasActions"]=&hasActions
	resp["allActions"]=&allActions
	resp["roleInfo"]=&roleInfo
	return resp
}
func SetRoleActiondb(idList []int,roleId int,resp map[string]interface{})map[string]interface{}{
	//查询到该角色
	var roleInfo=RoleInfo{}
	o:=orm.NewOrm()
	err:=o.QueryTable("role_info").Filter("del_flag",0).One(&roleInfo)
	if err!=nil{
		resp["flag"]="no"
		beego.Error(err)
		return resp
	}
	//查询到拥有的权限
	var hasActions=[]*ActionInfo{}
	 n,haserr:=o.QueryTable("action_info").Filter("Roles__RoleInfo__Id",roleId).Filter("" +
	 	"del_flag",0).All(&hasActions)
	 if haserr!=nil{
	 	if n==0{
	 		beego.Info("该角色没有权限")
		}else{
			resp["flag"]="no"
			beego.Error(err)
			return resp
		}
	 }
	 //清除已有的权限
	 o.Begin()
	 var remerr error
	 var adderr error
	 var queerr error
	 m2m:=o.QueryM2M(&roleInfo,"Actions")
	 for _,action:= range hasActions{
	 	 _,remerr=m2m.Remove(action)
	 }
	 //添加权限
	 var actionInfo=ActionInfo{}
	 for i:=0;i<len(idList);i++{
	 	queerr=o.QueryTable("action_info").Filter("id",idList[i]).One(&actionInfo)
	 	_,adderr=m2m.Add(actionInfo)
	 }
	 if remerr!=nil||adderr!=nil||queerr!=nil{
	 	o.Rollback()
	 }
	 o.Commit()
	resp["flag"]="ok"
	return resp

}