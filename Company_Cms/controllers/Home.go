package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func(this *HomeController)Index(){
	this.TplName="Home/Index.html"
}
func(this *HomeController)WindowIndex(){
	this.TplName="Home/WindowIndex.html"
}
