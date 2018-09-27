package controllers

import (
	"github.com/astaxie/beego"
	"Company_Cms/models"
	"strings"
	"strconv"
)

type RoleInfoController struct {
	beego.Controller
}

func(this *RoleInfoController)Index(){
	this.TplName="RoleInfo/Index.html"
}
func (this *RoleInfoController)GetRoleInfo(){
	//获取角色表的所有信息分页展示
   pageSize,_:=this.GetInt("rows")
   pageIndex,_:=this.GetInt(" page")

   var roles =[]models.RoleInfo{}
   resp=models.GetRoleInfodb(&roles, pageSize,pageIndex,resp)

   this.Data["json"]=resp
   this.ServeJSON()
}
func (this *RoleInfoController)ShowAddRole(){
	this.TplName="RoleInfo/ShowAddRole.html"
}
func(this *RoleInfoController)AddRole(){
	roleName:=this.GetString("roleName")
	roleRemark:=this.GetString("roleRemark")
	//添加至数据库
	resp=models.AddRoledb(roleName,roleRemark,resp)
	this.Data["json"]=resp
	this.ServeJSON()
}
func(this *RoleInfoController)DeleteRole(){
	strRoleId:=this.GetString("strRoleId")
	//将该角色删除
	resp=models.DeleteRoledb(strRoleId,resp)
	this.Data["json"]=resp
	this.ServeJSON()


}
func(this *RoleInfoController)ShowEditRole(){
	//获取iframe传过来的id
	id,_:=this.GetInt("id")
	//查询角色信息
	var roleInfo=models.RoleInfo{}
	resp=models.ShowEditRoledb(id,&roleInfo,resp)
	//将信息传回,将原有的角色信息填充
	this.Data["roleInfo"]=roleInfo
	//将视图显示
	this.TplName="RoleInfo/ShowEditRole.html"

}
func(this *RoleInfoController)EditRole(){
	editName:=this.GetString("editName")
	editRemark:=this.GetString("editRemark")
	editId,_:=this.GetInt("editId")
	resp=models.EditRoledb(editName,editRemark,editId,resp)
	//返回数据
	this.Data["json"]=resp
	this.ServeJSON()
}
func(this *RoleInfoController)ShowSetRoleAction(){
	//获取选中的角色id
	roleId,err:=this.GetInt("roleId")
	if err!=nil{
		resp["flag"]="no"
		beego.Error("获取角色id失败")
		return
	}
	//查询到该用户，将拥有角色信息和全部角色信息返回，
	resp=models.ShowSetRoleActiondb(roleId,resp)
	this.Data["roleInfo"]=resp["roleInfo"]
	this.Data["hasActions"]=resp["hasActions"]
	this.Data["allActions"]=resp["allActions"]

	this.TplName="RoleInfo/ShowSetRoleAction.html"
}
func(this *RoleInfoController)SetRoleAction(){
//获取要添加权限的角色id
	roleId,err:=this.GetInt("roleId")
	if err!=nil{
		resp["flag"]="no"
		beego.Error(err)
		return
	}
	//获取到选中的权限id
	var idList []int
	allkeys:=this.Ctx.Request.PostForm
	for key,_:=range allkeys{
		if strings.Contains(key,"cba_"){
			tempKey:=strings.Replace(key,"cba_","",-1)
			id,err:=strconv.Atoi(tempKey)
			if err!=nil{
				resp["flag"]="no"
				beego.Error(err)
				return
			}
			idList=append(idList,id)
		}
	}

	resp=models.SetRoleActiondb(idList,roleId,resp)
	this.Data["json"]=resp
	this.ServeJSON()
}