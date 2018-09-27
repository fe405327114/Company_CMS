package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

//权限信息
type ActionInfo struct {
	Id             int
	Remark         string
	DelFlag        int
	AddDate        time.Time
	ModifDate      time.Time
	Url            string
	HttpMethod     string
	ActionInfoName string
	ActionTypeEnum int
	MenuIcon       string
	IconWidth      int
	IconHeight     int
	Roles          []*RoleInfo   `orm:"reverse(many)"`
	UserAction     []*UserAction `orm:"reverse(many)"`
}

func GetActionInfodb(pageSize, pageNumber int, resp map[string]interface{}) map[string]interface{} {
	start := (pageNumber - 1) * pageSize
	//分页显示
	o := orm.NewOrm()
	var actionInfos = []ActionInfo{}
	n, err := o.QueryTable("action_info").Filter("del_flag", 0).OrderBy("id" +
		"").Limit(pageSize, start).All(&actionInfos)
	//获取总记录数
	count, err1 := o.QueryTable("action_info").Filter("del_flag", 0).Count()
	if err != nil || err1 != nil {
		if n == 0 {
			beego.Info("用户无权限信息")
		} else {
			resp["flag"] = "no"
			beego.Error("权限分页出错")
			return resp
		}
	}
	resp["flag"] = "ok"
	resp["rows"] = &actionInfos
	resp["total"] = count
	return resp
}
func DeleteActiondb(ids []int, resp map[string]interface{}) map[string]interface{} {
	o := orm.NewOrm()
	var actionInfo = ActionInfo{}
	for _, id := range ids {
		actionInfo.Id = id
		_, err := o.Delete(&actionInfo)
		if err != nil {
			resp["flag"] = "no"
			return resp
		}
	}
	beego.Info(ids)
	resp["flag"] = "ok"
	return resp
}
func ShowEditActiondb(actionId int, resp map[string]interface{}) map[string]interface{} {
	//查询该权限，返回视图
	o := orm.NewOrm()
	var actionInfo = ActionInfo{}
	erroA := o.QueryTable("action_info").Filter("id", actionId).Filter("del_flag"+
		"", 0).One(&actionInfo)
	if erroA != nil {
		resp["flag"] = "no"
		beego.Error(erroA)
		return resp
	}
	resp["actionInfo"] = &actionInfo
	return resp
}
