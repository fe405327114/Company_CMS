package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"Company_Cms/models"
)

type LoginController struct {
	beego.Controller
}

func(this *LoginController)Index(){
	this.TplName="Login/Index.html"
}
func(this *LoginController)UserLogin(){
	userName:=this.GetString("LoginCode")
	userPwd:=this.GetString("LoginPwd")
	o:=orm.NewOrm()
	var userInfo=models.UserInfo{}
	err:=o.QueryTable("user_info").Filter("user_name",userName).Filter("user_pwd" +
		"",userPwd).One(&userInfo)
	if err!=nil{
		resp["flag"]="no"
		beego.Error(err)
		return
	}
	if userInfo.Id>0{
		this.SetSession("userId",userInfo.Id)
		this.SetSession("userName",userInfo.UserName)
		resp["flag"]="ok"
	}else {
		resp["flag"]="no"
	}
    this.Data["json"]=resp
	this.ServeJSON()
}
