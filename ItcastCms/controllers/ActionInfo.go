package controllers

import (
	"github.com/astaxie/beego"
	"ItcastCms/models"
	"path"
	"time"
	"strconv"
	"os"
	"github.com/astaxie/beego/orm"
)

type ActionInfoController struct {
	beego.Controller
}
func(this *ActionInfoController)Index(){
	//显示视图
	this.TplName="ActionInfo/Index.html"
}
func(this *ActionInfoController)GetActionInfo(){
	//获取datagrid中的分页数据
	pageSize,err:=this.GetInt("rows")
	pageNumber,err1:=this.GetInt("page")
	if err!=nil||err1!=nil{
		resp["flag"]="no"
		beego.Error("获取分页数据失败!",err,err1)
		return
	}
	resp=models.GetActionInfodb(pageSize,pageNumber,resp)
	this.Data["json"]=resp
	this.ServeJSON()
}
func(this *ActionInfoController)ShowAddAction(){
	//将iframe试图返回给它
	this.TplName="ActionInfo/Index.html"
}
func(this *ActionInfoController)FileUp(){
	f,h,err:=this.GetFile("fileUp")
	if err!=nil{
		resp["flag"]="no"
		resp["msg"]="文件上传失败"
		return
	}
	defer f.Close()
	//判断图片信息
	if h.Size>500000{
		resp["flag"]="no"
		resp["msg"]="文件上传过大"
		return
	}
      beego.Error("33333333333333")
	fileExt:=path.Ext(h.Filename)
	if fileExt!=".png"&&fileExt!=".jpg"{
		resp["flag"]="no"
		resp["msg"]="文件类型错误"
		return
	}
	beego.Error("44444444444")
	//创建文件夹，以日期命名
	dir:="./static/fileUp/"+strconv.Itoa(time.Now().Year())+"-"+time.Now().Month().String()+"-"+
		strconv.Itoa(time.Now().Day())+"/"
	//判断路径是否存在
	_,err=os.Stat(dir)
	if err!=nil{
		//创建文件夹
		err1:=os.Mkdir(dir,os.ModePerm)
		if err1!=nil{
			resp["flag"]="no"
			resp["msg"]="服务器创建文件夹错误"
			return
		}
	}

	//time.Parse("2006-01-02 15-04-05",参数)将字符串转换为时间
	fileName:=time.Now().Format("2006-01-02 15-04-05")
   beego.Error("111111111111")
 	errsav:=this.SaveToFile("fileUp",dir+fileName+fileExt)
	if errsav!=nil{
		resp["flag"]="no"
		resp["msg"]="服务器保存文件夹错误"
		beego.Error(errsav)
		return
	}
	beego.Error("222222222")
	resp["flag"]="ok"
	resp["msg"]=dir+fileName+fileExt
	this.Data["json"]=resp
	this.ServeJSON()
}
func(this *ActionInfoController)AddAction(){
	var actionInfo=models.ActionInfo{}
	actionInfo.ActionTypeEnum,_=this.GetInt("ActionTypeEnum")
	actionInfo.MenuIcon=this.GetString("MenuIcon")
	actionInfo.Url=this.GetString("Url")
	actionInfo.ActionInfoName=this.GetString("ActionInfoName")
	actionInfo.IconWidth=0
	actionInfo.IconHeight=0
	actionInfo.HttpMethod=this.GetString("HttpMethod")
	actionInfo.Remark=this.GetString("Remark")
	actionInfo.DelFlag=0
	actionInfo.AddDate=time.Now()
	actionInfo.ModifDate=time.Now()
	o:=orm.NewOrm()
	_,err:=o.Insert(&actionInfo)
	if err!=nil{
		resp["flag"]="no"
		beego.Error("插入权限信息失败",err)
		return
	}
	resp["flag"]="ok"
	this.Data["json"]=resp
	this.ServeJSON()
}
