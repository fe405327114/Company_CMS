package controllers

import (
	"github.com/astaxie/beego"
	"Company_Cms/models"
)

type HomeController struct {
	beego.Controller
}

func(this *HomeController)Index(){
	this.TplName="Home/Index.html"
}
func(this *HomeController)WindowIndex(){
	this.TplName="Home/WindowIndex.html"
}
func(this *HomeController)MenuSelect(){
	//第一：获取登录用户的信息
	//第二：根据登录用户，查询出该登录用户具有的角色信息。
	//第三：根据角色信息，查询出对应的权限信息。
	//第四：对查询出的权限信息进行判断，判断一下是否为菜单权限，也就是ActionTypeEnum属性的取值是否为1.
	//第五：上面是按照“用户--角色--权限”，这条线进行权限过滤的，那么下面就是按照“用户--权限""这条线进行过滤。
	//第六：按照“用户---权限”这条线查询出用户对应的权限后，还要进行菜单权限的过滤。
	//第七：将两条线查询出的权限进行合并。
	//第八：去重操作
	//第九：过滤掉登录用户禁用的权限。也就是判断登录用户中关于isPass的取值。
	//第十：前端内容的修改。
		userId:=this.GetSession("userId")

		//开始菜单权限过滤
	    resp= models.MenuSelectdb(userId.(int),resp)

	    this.Data["json"]=resp
	    this.ServeJSON()
	}

