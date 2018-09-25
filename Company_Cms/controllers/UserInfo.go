package controllers

import (
	"github.com/astaxie/beego"
    "ItcastCms/models"
	"time"
	_"github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"github.com/astaxie/beego/orm"
	"ItcastCms/viewmodels"
)

type UserInfoController struct {
	beego.Controller
}

func(this *UserInfoController)Index(){
	this.TplName="UserInfo/Index.html"
}
var resp =make(map[string]interface{})
func(this *UserInfoController)AddUser(){
	userName:=this.GetString("userName")
	userPwd:=this.GetString("userPwd")
	userRemark:=this.GetString("userRemark")

	var userInfo=models.UserInfo{}
	userInfo.Remark=userRemark
	userInfo.UserName=userName
	userInfo.UserPwd=userPwd
	userInfo.AddDate=time.Now()
	userInfo.ModifDate=time.Now()
	userInfo.DelFlag=0 //0表示没删除
	//插入数据库

	resp:=models.AddUserdb(&userInfo,resp)

    this.Data["json"]= resp
    this.ServeJSON()
}
func (this *UserInfoController)GetUserInfo(){
	/*qb, _ := orm.NewQueryBuilder("mysql")
	//分页数据
	qb.Select("*").From("user_info").Where("del_flag=?").OrderBy(
		"id").Asc().Limit(pageSize).Offset(start)
	sql := qb.String()
	o := orm.NewOrm()
	//o.QueryTable("user_info").Filter("del_flag",0).OrderBy(
	//	"-id").Limit(pageSize,start).All(userInfos)
	_, err := o.Raw(sql, 0).QueryRows(userInfos)
	count,_:=o.QueryTable("user_info").Filter("del_flag",0).Count()
	if err != nil {
		resp["flag"] = "no"
	} else {
		resp["flag"] = "ok"
		resp["rows"] = userInfos
		resp["total"]=count
	}*/

	//接收前端发送的当前页码与每页显示的记录数
	//前端传的数量和当前页面的ke分别是page和rows
	pageIndex,_:=this.GetInt("page")
	pageSize,_:=this.GetInt("rows")

	name:=this.GetString("name")
	remark:=this.GetString("remark")
	//给查询对象赋值
	var userData=viewmodels.UserData{}
	userData.Name=name
	userData.Remark=remark
	userData.PageSize=pageSize
	userData.PageIndex=pageIndex
	start:=(pageIndex-1)*pageSize
	userData.Start=start
	//查询数据
	userInfos:=userData.SearchUserData()
	resp["rows"]=userInfos
	resp["total"]=userData.TotalCount
	this.Data["json"]=resp
	//this.Data["json"]= resp
	this.ServeJSON()
}
func(this *UserInfoController)DeleteUser(){
  //接受传递的值
  ids:=this.GetString("strId")
  strId:=strings.Split(ids,",")
  var userInfo=models.UserInfo{}
	resp:=models.DeleteUserdb(&userInfo,strId,resp)
   this.Data["json"]=resp
   this.ServeJSON()

}
func (this *UserInfoController)ShowEditUser(){
	id,_:=this.GetInt("id")
	beego.Info("id",id)
	//查询该数据并返回
	var userInfo=models.UserInfo{}
	userInfo.Id=id
	resp:=models.ShowEditUserdb(&userInfo,resp)

	this.Data["json"]=resp
	beego.Info("resp",resp)
	this.ServeJSON()
}
func(this *UserInfoController)EditUser(){
	//获取输入的内容
	editName:=this.GetString("editName")
	editPwd:=this.GetString("editPwd")
	editRemark:=this.GetString("editRemark")
	Eidtid:=this.GetString("Eidtid")
	EditaddDate:=this.GetString("EditaddDate")
	//进行更新操作
	 var userInfo=models.UserInfo{}
	userInfo.UserName=editName
	userInfo.UserPwd=editPwd
	userInfo.Remark=editRemark
	userInfo.ModifDate=time.Now()
	id,_:=strconv.Atoi(Eidtid)
	userInfo.Id=id
	//format将时间转换成字符串 ，parse将字符串转换成时间
	//注意这里的时间和表格里的时间格式联系
	t,err0:=time.Parse("2006-01-02T15:04:05+08:00",EditaddDate)
	if err0!=nil{
		beego.Error(err0)
		return
	}
	userInfo.AddDate=t
	userInfo.DelFlag=0
	resp:=models.EditUserdb(&userInfo,resp)
	this.Data["json"]=resp
	this.ServeJSON()
}
func(this *UserInfoController)AddUserRole(){
	//获取用户id，返回视图
	id,_:=this.GetInt("id")
	o:=orm.NewOrm()
	//查询选中的用户
	var userInfo=models.UserInfo{}
	err:=o.QueryTable("user_info").Filter("id",id).One(&userInfo)
	//查询该用户拥有的角色和所有的角色，在视图中循环显示
	var roles=[]models.RoleInfo{}
	_,err2:=o.QueryTable("role_info").Filter("del_flag",0).All(&roles)
	var hasRoles=[]*models.RoleInfo{}
     //多对多查询
	n,err1:=o.QueryTable("role_info").Filter("Users__UserInfo__Id",id).
		Filter("del_flag",0).All(&hasRoles)
	//如果用户和角色是一对多关系，则用以下方法查询
	//o.QueryTable("role_info").RelatedSel("user_info").Filter(
	//	"UserInfo_Id",id).All(&hasRoles)
	if err==nil&&err1==nil&&err2==nil{
		//模板中需要用到userInfo中的Id和Name
		this.Data["userName"]=&userInfo
        this.Data["roles"]=&roles
        this.Data["hasRoles"]=&hasRoles
	}else if n==0{
			this.Data["userName"]=&userInfo
			this.Data["roles"]=&roles
			this.Data["hasRoles"]=&hasRoles
		}else {
			beego.Error(err)
			return
		}
	this.TplName="UserInfo/AddUserRole.html"
	}
func(this *UserInfoController)EditRole(){
        //获取选中的角色id信息
        userId,err:=this.GetInt("userId")
        if err!=nil{
        	resp["flag"]="no"
        	beego.Error("获取编辑角色的用户id失败！",err)
        	return
		}
        //获取所有以_che开头的数据
        allkeys:=this.Ctx.Request.PostForm  //url.Values类型
        var idList []int
        for key,_:= range allkeys{
			if strings.Contains(key,"che_"){
				//把che_去掉
				strId:=strings.Replace(key,"che_","",-1)
				id,_:=strconv.Atoi(strId)
				idList=append(idList,id)
			}
		}
		resp=models.EditUserRoledb(userId,idList,resp)
		this.Data["json"]=resp
		this.ServeJSON()
}
func(this *UserInfoController)ShowUserAction(){
	//获取iframe  get发送的选中的用户id
	userId,errget:=this.GetInt("id")
	if errget!=nil{
		resp["flag"]="no"
		beego.Error(errget)
		return
	}
	resp=models.ShowUserActiondb(userId,resp)
	resp["flag"]="ok"
	this.Data["userInfo"]=resp["userInfo"]
	this.Data["userExtActions"]=resp["userExtActions"]
	this.Data["allActions"]=resp["allActions"]
	this.TplName="UserInfo/ShowUserAction.html"
}
func(this*UserInfoController)SetUserAction(){
	userId,erruI:=this.GetInt("userId")
	actionId,erraI:=this.GetInt("actionId")
	Ispass,erraTE:=this.GetInt("Ispass")
	if erruI!=nil||erraI!=nil||erraTE!=nil{
		resp["flag"]="no"
		beego.Error(erruI,erraI,erraTE)
		return
	}
	resp=models.SetUserActiondb(userId,actionId,Ispass,resp)
	this.Data["json"]=resp
	this.ServeJSON()
}
func(this *UserInfoController)DeleteUserAction(){
	//获取要删除的用户和权限id
	userId,erruI:=this.GetInt("userId")
	actionId,erraI:=this.GetInt("actionId")
	if erraI!=nil||erruI!=nil{
		resp["flag"]="no"
		return
	}
	resp=models.DeleteUserActiondb(userId,actionId,resp)
	this.Data["json"]=resp
	this.ServeJSON()
}

