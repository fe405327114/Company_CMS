package controllers

import (
	"Company_Cms/models"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type ArticelClassController struct {
	beego.Controller
}

func(this *ArticelClassController) Index()  {
this.TplName="ArticelClass/Index.html"
}
//展示添加根类别的表单页面。
func (this *ArticelClassController)ShowAddParent()  {
  this.TplName="ArticelClass/ShowAddParent.html"
}
//添加根类别。
func (this *ArticelClassController)AddParentClass()  {
	var articelClass=models.ArticelClass{}
	articelClass.DelFlag=0
	articelClass.ClassName=this.GetString("className")
	articelClass.Remark=this.GetString("remark")
	articelClass.ParentId=0 //表示根
	articelClass.CreateDate=time.Now()
	articelClass.CreateUserId=this.GetSession("userId").(int)
	o:=orm.NewOrm()
	_,err:=o.Insert(&articelClass)
	if err==nil{
		this.Data["json"]=map[string]interface{}{"flag":"ok"}
	}else{
		this.Data["json"]=map[string]interface{}{"flag":"no"}
	}
	this.ServeJSON()
}

//查询根类别信息
func (this *ArticelClassController)ShowArticelClass()  {
  o:=orm.NewOrm()
  var articelClass []models.ArticelClass
  o.QueryTable("articel_class").Filter("parent_id",0).All(&articelClass)
  this.Data["json"]=map[string]interface{}{"rows":articelClass}//返回的数据赋值给rows
  this.ServeJSON()


}
//展示添加子类别的页面。
func (this *ArticelClassController)ShowAddChildClass()  {
	cId,_:=this.GetInt("cId")
	//查询出对应的根类别信息。
	o:=orm.NewOrm()
	var articelClass models.ArticelClass
	o.QueryTable("articel_class").Filter("id",cId).One(&articelClass)
	this.Data["classInfo"]=articelClass
	this.TplName="ArticelClass/ShowAddChildClass.html"
}
//完成子类别的添加
func (this *ArticelClassController)AddChildClass()  {
  var classInfo=models.ArticelClass{}
  classInfo.CreateUserId=this.GetSession("userId").(int)
  classInfo.CreateDate=time.Now()
  classInfo.ParentId,_=this.GetInt("classId")
	classInfo.Remark=this.GetString("remark")
	classInfo.ClassName=this.GetString("className")
	o:=orm.NewOrm()
	_,err:=o.Insert(&classInfo)
	if err==nil{
		this.Data["json"]=map[string]interface{}{"flag":"ok"}
	}else{
		this.Data["json"]=map[string]interface{}{"flag":"no"}
	}
	this.ServeJSON()

}
//查询根下面的子类别信息。
func (this *ArticelClassController)ShowChildClass()  {
	cid,_:=this.GetInt("id")//根的编号。
	o:=orm.NewOrm()
	var articelClasses[]models.ArticelClass
	o.QueryTable("articel_class").Filter("parent_id",cid).All(&articelClasses)
	this.Data["json"]=map[string]interface{}{"rows":articelClasses}
	this.ServeJSON()
}